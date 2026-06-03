#!/usr/bin/env python3
# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os
import re
import glob
import json
import subprocess
import argparse
from datetime import datetime, timezone
from http.server import HTTPServer, BaseHTTPRequestHandler

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
REPO_ROOT = os.path.abspath(os.path.join(SCRIPT_DIR, "..", "..", "..", ".."))
CHECKLIST_DIR = os.path.join(REPO_ROOT, ".gemini", "skills", "checklist-for-kind")
DB_DIR = os.path.join(CHECKLIST_DIR, "database")
OUTPUT_REPO_PATH = os.path.join(CHECKLIST_DIR, "journal", "migration_dashboard.html")
OUTPUT_ARTIFACT_PATH = "/usr/local/google/home/justinsb/.gemini/jetski/html_artifacts/migration_dashboard.html"

HTML_TEMPLATE = """<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>KCC Direct Migration Dashboard</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;500;600;700&family=JetBrains+Mono:wght@400;500&display=swap" rel="stylesheet">
    <style>
        :root {{
            --bg-color: #0f1319;
            --panel-bg: rgba(22, 28, 38, 0.6);
            --panel-border: rgba(255, 255, 255, 0.08);
            --text-primary: #f3f4f6;
            --text-secondary: #9ca3af;
            
            --status-completed: #10b981;
            --status-completed-bg: rgba(16, 185, 129, 0.15);
            --status-active: #0ea5e9;
            --status-active-bg: rgba(14, 165, 233, 0.15);
            --status-open: #f59e0b;
            --status-open-bg: rgba(245, 158, 11, 0.15);
            --status-pending: #6b7280;
            --status-pending-bg: rgba(107, 114, 128, 0.1);
            
            --font-main: 'Outfit', sans-serif;
            --font-code: 'JetBrains Mono', monospace;
            --transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        }}

        * {{
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }}

        body {{
            background-color: var(--bg-color);
            color: var(--text-primary);
            font-family: var(--font-main);
            padding: 2.5rem;
            min-height: 100vh;
            line-height: 1.5;
            background-image: radial-gradient(circle at 10% 20%, rgba(14, 165, 233, 0.05) 0%, transparent 40%),
                              radial-gradient(circle at 90% 80%, rgba(16, 185, 129, 0.03) 0%, transparent 40%);
        }}

        .container {{
            max-width: 1200px;
            margin: 0 auto;
        }}

        header {{
            margin-bottom: 3rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            border-bottom: 1px solid var(--panel-border);
            padding-bottom: 1.5rem;
        }}

        h1 {{
            font-size: 2.25rem;
            font-weight: 700;
            letter-spacing: -0.025em;
            background: linear-gradient(135deg, #ffffff 0%, #a5b4fc 100%);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }}

        .last-updated {{
            font-size: 0.875rem;
            color: var(--text-secondary);
            font-family: var(--font-code);
        }}

        /* Stats Grid */
        .stats-grid {{
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
            gap: 1.5rem;
            margin-bottom: 3rem;
        }}

        .stat-card {{
            background: var(--panel-bg);
            border: 1px solid var(--panel-border);
            border-radius: 16px;
            padding: 1.5rem;
            backdrop-filter: blur(12px);
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            transition: var(--transition);
        }}

        .stat-card:hover {{
            transform: translateY(-2px);
            border-color: rgba(255, 255, 255, 0.15);
            box-shadow: 0 10px 20px -10px rgba(0, 0, 0, 0.5);
        }}

        .stat-label {{
            font-size: 0.875rem;
            font-weight: 500;
            color: var(--text-secondary);
            margin-bottom: 0.5rem;
            text-transform: uppercase;
            letter-spacing: 0.05em;
        }}

        .stat-value {{
            font-size: 2rem;
            font-weight: 700;
            color: #ffffff;
        }}

        /* Kind Section Card */
        .kind-card {{
            background: var(--panel-bg);
            border: 1px solid var(--panel-border);
            border-radius: 20px;
            padding: 2rem;
            margin-bottom: 2rem;
            backdrop-filter: blur(12px);
            transition: var(--transition);
        }}

        .kind-card:hover {{
            border-color: rgba(255, 255, 255, 0.15);
            box-shadow: 0 12px 30px -15px rgba(0, 0, 0, 0.6);
        }}

        .kind-header {{
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 1.5rem;
        }}

        .kind-title {{
            display: flex;
            align-items: center;
            gap: 1rem;
        }}

        .kind-name {{
            font-size: 1.5rem;
            font-weight: 600;
            color: #ffffff;
        }}

        .kind-badge {{
            font-size: 0.75rem;
            font-weight: 600;
            text-transform: uppercase;
            letter-spacing: 0.05em;
            padding: 0.25rem 0.75rem;
            border-radius: 20px;
        }}

        /* Progress Bar */
        .progress-container {{
            display: flex;
            align-items: center;
            gap: 1rem;
            width: 300px;
        }}

        .progress-bar-bg {{
            flex-grow: 1;
            height: 8px;
            background: rgba(255, 255, 255, 0.05);
            border-radius: 4px;
            overflow: hidden;
        }}

        .progress-bar-fill {{
            height: 100%;
            background: linear-gradient(90deg, var(--status-active) 0%, var(--status-completed) 100%);
            border-radius: 4px;
            transition: width 1s ease-in-out;
        }}

        .progress-text {{
            font-size: 0.875rem;
            font-weight: 600;
            color: var(--text-secondary);
            font-family: var(--font-code);
            white-space: nowrap;
        }}

        /* Steps Timeline */
        .steps-container {{
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 1rem;
            position: relative;
            margin-top: 1.5rem;
        }}

        .step-node {{
            background: rgba(255, 255, 255, 0.02);
            border: 1px solid var(--panel-border);
            border-radius: 12px;
            padding: 1.25rem;
            display: flex;
            flex-direction: column;
            gap: 0.75rem;
            position: relative;
            transition: var(--transition);
        }}

        .step-node:hover {{
            background: rgba(255, 255, 255, 0.04);
            border-color: rgba(255, 255, 255, 0.12);
        }}

        .step-header {{
            display: flex;
            justify-content: space-between;
            align-items: center;
        }}

        .step-number {{
            font-size: 0.75rem;
            font-weight: 700;
            color: var(--text-secondary);
            font-family: var(--font-code);
        }}

        .step-status-indicator {{
            width: 8px;
            height: 8px;
            border-radius: 50%;
        }}

        .step-name {{
            font-size: 0.95rem;
            font-weight: 500;
            color: #ffffff;
        }}

        .step-links {{
            display: flex;
            flex-wrap: wrap;
            gap: 0.5rem;
            margin-top: auto;
        }}

        .link-badge {{
            font-size: 0.75rem;
            font-family: var(--font-code);
            text-decoration: none;
            padding: 0.15rem 0.5rem;
            border-radius: 6px;
            transition: var(--transition);
        }}

        /* Custom Status Colors */
        .state-completed {{
            border-color: rgba(16, 185, 129, 0.3);
            background: rgba(16, 185, 129, 0.03);
        }}
        .state-completed .step-status-indicator {{
            background-color: var(--status-completed);
            box-shadow: 0 0 8px var(--status-completed);
        }}
        .state-completed .link-badge {{
            background: var(--status-completed-bg);
            color: var(--status-completed);
        }}
        .state-completed .link-badge:hover {{
            background: rgba(16, 185, 129, 0.25);
        }}

        .state-active {{
            border-color: rgba(14, 165, 233, 0.4);
            background: rgba(14, 165, 233, 0.03);
        }}
        .state-active .step-status-indicator {{
            background-color: var(--status-active);
            box-shadow: 0 0 8px var(--status-active);
        }}
        .state-active .link-badge {{
            background: var(--status-active-bg);
            color: var(--status-active);
        }}
        .state-active .link-badge:hover {{
            background: rgba(14, 165, 233, 0.25);
        }}

        .state-open {{
            border-color: rgba(245, 158, 11, 0.4);
            background: rgba(245, 158, 11, 0.03);
        }}
        .state-open .step-status-indicator {{
            background-color: var(--status-open);
            box-shadow: 0 0 8px var(--status-open);
        }}
        .state-open .link-badge {{
            background: var(--status-open-bg);
            color: var(--status-open);
        }}
        .state-open .link-badge:hover {{
            background: rgba(245, 158, 11, 0.25);
        }}

        .state-pending {{
            opacity: 0.6;
        }}
        .state-pending .step-status-indicator {{
            background-color: var(--status-pending);
        }}

        @media (max-width: 900px) {{
            .steps-container {{
                grid-template-columns: 1fr 1fr;
            }}
        }}
        @media (max-width: 600px) {{
            .steps-container {{
                grid-template-columns: 1fr;
            }}
            header {{
                flex-direction: column;
                align-items: flex-start;
                gap: 1rem;
            }}
        }}
        #sync-btn:hover {{
            background: rgba(14, 165, 233, 0.25);
            border-color: rgba(14, 165, 233, 0.6);
            box-shadow: 0 0 10px rgba(14, 165, 233, 0.3);
        }}
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div>
                <h1>KCC Direct Migration Pipeline</h1>
            </div>
            <div style="display: flex; align-items: center; gap: 1rem;">
                <div id="sync-btn-container" style="display: none;">
                    <button onclick="triggerSync()" id="sync-btn" style="background: var(--status-active-bg); color: var(--status-active); border: 1px solid rgba(14, 165, 233, 0.3); padding: 0.35rem 1rem; border-radius: 8px; font-family: var(--font-main); font-weight: 500; cursor: pointer; transition: var(--transition);">
                        Sync Status
                    </button>
                </div>
                <div class="last-updated">Last Sync: {sync_time} UTC</div>
            </div>
        </header>

        <!-- General Status -->
        <div class="stats-grid">
            <div class="stat-card">
                <div class="stat-label">Active Resources</div>
                <div class="stat-value">{total_resources}</div>
            </div>
            <div class="stat-card">
                <div class="stat-label">Total Steps Completed</div>
                <div class="stat-value">{completed_steps} / {total_steps}</div>
            </div>
            <div class="stat-card">
                <div class="stat-label">PRs Staged / Open</div>
                <div class="stat-value">{open_prs}</div>
            </div>
            <div class="stat-card">
                <div class="stat-label">Fully Reconciled</div>
                <div class="stat-value">{fully_reconciled}</div>
            </div>
        </div>

        {cards_html}

    </div>
    <script>
        if (window.location.protocol.startsWith('http')) {{
            document.getElementById('sync-btn-container').style.display = 'block';
        }}
        function triggerSync() {{
            const btn = document.getElementById('sync-btn');
            btn.disabled = true;
            btn.innerText = 'Syncing...';
            btn.style.opacity = '0.7';
            fetch('/sync', {{ method: 'POST' }})
                .then(response => {{
                    if (response.ok) {{
                        window.location.reload();
                    }} else {{
                        alert('Sync failed!');
                        btn.disabled = false;
                        btn.innerText = 'Sync Status';
                        btn.style.opacity = '1';
                    }}
                }})
                .catch(err => {{
                    alert('Error: ' + err);
                    btn.disabled = false;
                    btn.innerText = 'Sync Status';
                    btn.style.opacity = '1';
                }});
        }}
    </script>
</body>
</html>
"""

def parse_yaml(content):
    """Parse nested dict from a simple YAML string schema for database config."""
    data = {}
    current_step_num = None
    for line in content.splitlines():
        line = line.rstrip()
        if not line or line.startswith("#"):
            continue
        if not line.startswith(" ") and ":" in line:
            k, v = line.split(":", 1)
            key = k.strip()
            val = v.strip().strip('"')
            if key == "steps":
                continue
            data[key] = val
            # convert current_step to int
            if key == "current_step":
                data[key] = int(val)
        elif line.startswith("  ") and not line.startswith("    ") and ":" in line:
            k, _ = line.split(":", 1)
            current_step_num = int(k.strip())
            if "steps" not in data:
                data["steps"] = {}
            data["steps"][current_step_num] = {}
        elif line.startswith("    ") and ":" in line:
            k, v = line.split(":", 1)
            raw_v = v.strip().strip('"')
            # Parse ints for issue and pr
            if k.strip() in ["issue", "pr"]:
                raw_v = int(raw_v) if raw_v.isdigit() else 0
            data["steps"][current_step_num][k.strip()] = raw_v
    return data

def write_yaml(data, filepath):
    """Serializes data dictionary back to a structured YAML file."""
    content = []
    content.append(f"kind: {data['kind']}")
    content.append(f"current_step: {data['current_step']}")
    content.append("steps:")
    for step_num in sorted(data["steps"].keys()):
        step = data["steps"][step_num]
        content.append(f"  {step_num}:")
        content.append(f'    name: "{step["name"]}"')
        content.append(f'    issue: {step.get("issue", 0)}')
        content.append(f'    pr: {step.get("pr", 0)}')
        content.append(f'    status: "{step.get("status", "Not Started")}"')
        content.append(f'    started: "{step.get("started", "")}"')
        content.append(f'    completed: "{step.get("completed", "")}"')
    with open(filepath, "w") as f:
        f.write("\n".join(content) + "\n")

def fetch_github_issue_state(issue_number):
    """Fetches details of a GitHub issue including state and linked pull requests."""
    cmd = [
        "gh", "issue", "view", str(issue_number),
        "-R", "GoogleCloudPlatform/k8s-config-connector",
        "--json", "state,closedByPullRequestsReferences"
    ]
    try:
        result = subprocess.run(cmd, capture_output=True, text=True, check=True)
        return json.loads(result.stdout)
    except Exception as e:
        print(f"Error fetching issue #{issue_number}: {e}")
        return None

def sync_database_record_with_github(filepath):
    """Refreshes step statuses and PR link references from GitHub API."""
    with open(filepath, "r") as f:
        data = parse_yaml(f.read())
        
    updated = False
    today_str = datetime.now(timezone.utc).strftime("%Y-%m-%d")
    
    for step_num, step in data.get("steps", {}).items():
        issue_num = step.get("issue", 0)
        if issue_num > 0:
            gh_issue = fetch_github_issue_state(issue_num)
            if gh_issue:
                # Get state
                state = gh_issue.get("state", "OPEN")
                
                # Get linked PR
                pr_num = 0
                pr_refs = gh_issue.get("closedByPullRequestsReferences", [])
                if pr_refs:
                    pr_num = pr_refs[0].get("number", 0)
                
                # Sync PR
                if step.get("pr", 0) != pr_num:
                    step["pr"] = pr_num
                    updated = True
                
                # Sync status
                status = "Not Started"
                if state == "CLOSED":
                    status = "Completed"
                else: # OPEN
                    status = "PR Created" if pr_num > 0 else "Open"
                    
                if step.get("status") != status:
                    step["status"] = status
                    updated = True
                    
                # Sync started / completed dates
                if status in ["Open", "PR Created", "Completed"] and not step.get("started"):
                    step["started"] = today_str
                    updated = True
                if status == "Completed" and not step.get("completed"):
                    step["completed"] = today_str
                    updated = True
                    
    # Sync current_step field
    current_step = 1
    for step_num in sorted(data["steps"].keys()):
        status = data["steps"][step_num]["status"]
        if status in ["Open", "PR Created"]:
            current_step = step_num
            break
        elif status == "Completed":
            current_step = step_num + 1 # Next step pending
            
    # Clip current step to max steps count (usually 4)
    max_steps = len(data["steps"])
    if current_step > max_steps:
        current_step = max_steps
        
    if data["current_step"] != current_step:
        data["current_step"] = current_step
        updated = True

    if updated:
        write_yaml(data, filepath)
        print(f"Synced and updated database record: {filepath}")
        
    return data

def build_step_html(step_num, step_data):
    """Generate HTML for a single step node."""
    status_class = "state-pending"
    status_lbl = "Planned"
    
    status = step_data["status"]
    if status == "Completed":
        status_class = "state-completed"
        status_lbl = "Completed"
    elif status == "PR Created":
        status_class = "state-active"
        status_lbl = "PR Active"
    elif status == "Open":
        status_class = "state-open"
        status_lbl = "Issue Open"

    links_html = []
    
    issue_num = step_data.get("issue", 0)
    if issue_num > 0:
        links_html.append(
            f'<a href="https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/{issue_num}" class="link-badge" target="_blank">#{issue_num}</a>'
        )

    pr_num = step_data.get("pr", 0)
    if pr_num > 0:
        links_html.append(
            f'<a href="https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/{pr_num}" class="link-badge" target="_blank">#{pr_num}</a>'
        )

    if not links_html:
        links_html.append(f'<span class="link-badge" style="background: transparent; border: 1px dashed rgba(255,255,255,0.1); color: var(--text-secondary);">{status_lbl}</span>')

    links_str = "\n".join(links_html)

    return f"""
                <div class="step-node {status_class}">
                    <div class="step-header">
                        <span class="step-number">STEP {step_num}</span>
                        <div class="step-status-indicator"></div>
                    </div>
                    <span class="step-name">{step_data["name"]}</span>
                    <div class="step-links">
                        {links_str}
                    </div>
                </div>"""

def build_card_html(record):
    """Generate HTML card for a resource kind."""
    total = len(record["steps"])
    completed = sum(1 for step_num in record["steps"] if record["steps"][step_num]["status"] == "Completed")
    
    progress_pct = (completed / total * 100) if total > 0 else 0
    
    # Determine current step badge labels and styles
    badge_style = 'background: rgba(255, 255, 255, 0.05); color: var(--text-secondary);'
    badge_lbl = "Not Started"
    
    active_found = False
    for step_num in sorted(record["steps"].keys()):
        status = record["steps"][step_num]["status"]
        if status in ["Open", "PR Created"]:
            active_found = True
            if status == "PR Created":
                badge_style = 'background: var(--status-active-bg); color: var(--status-active);'
                badge_lbl = f"Step {step_num} PR"
            else:
                badge_style = 'background: var(--status-open-bg); color: var(--status-open);'
                badge_lbl = f"Step {step_num} Open"
            break
            
    if not active_found and completed == total:
        badge_style = 'background: var(--status-completed-bg); color: var(--status-completed);'
        badge_lbl = "Fully Reconciled"
    elif not active_found and completed > 0:
        # Steps are done, but next step is not started/active
        badge_style = 'background: rgba(255, 255, 255, 0.05); color: var(--text-secondary);'
        badge_lbl = f"Step {completed + 1} Pending"

    steps_html = "".join(build_step_html(step_num, record["steps"][step_num]) for step_num in sorted(record["steps"].keys()))
    
    # Pad empty column if only 3 steps exist
    if total == 3:
        steps_html += '\n                <div class="step-node state-pending" style="visibility: hidden;"></div>'

    return f"""
        <!-- {record["kind"]} -->
        <div class="kind-card">
            <div class="kind-header">
                <div class="kind-title">
                    <span class="kind-name">{record["kind"]}</span>
                    <span class="kind-badge" style="{badge_style}">{badge_lbl}</span>
                </div>
                <div class="progress-container">
                    <div class="progress-bar-bg">
                        <div class="progress-bar-fill" style="width: {progress_pct}%;"></div>
                    </div>
                    <span class="progress-text">{completed} / {total} Steps</span>
                </div>
            </div>
            <div class="steps-container">
                {steps_html}
            </div>
        </div>"""

def load_database_records():
    db_records = glob.glob(os.path.join(DB_DIR, "*.yaml"))
    records = []
    for filepath in db_records:
        with open(filepath, "r") as f:
            rec = parse_yaml(f.read())
            records.append(rec)
    records.sort(key=lambda x: x["kind"])
    return records

def sync_and_generate():
    db_records = glob.glob(os.path.join(DB_DIR, "*.yaml"))
    for filepath in db_records:
        sync_database_record_with_github(filepath)
    generate_dashboard_html()

def generate_dashboard_html():
    records = load_database_records()
    
    # Global stats
    total_resources = len(records)
    completed_steps = 0
    total_steps = 0
    open_prs = 0
    fully_reconciled = 0

    cards_list = []
    for r in records:
        total = len(r["steps"])
        completed = sum(1 for step_num in r["steps"] if r["steps"][step_num]["status"] == "Completed")
        
        total_steps += total
        completed_steps += completed
        open_prs += sum(1 for step_num in r["steps"] if r["steps"][step_num]["status"] == "PR Created")
        if completed == total and total > 0:
            fully_reconciled += 1
            
        cards_list.append(build_card_html(r))
        
    cards_html = "\n".join(cards_list)
    sync_time = datetime.now(timezone.utc).strftime("%Y-%m-%d %H:%M:%S")

    html_out = HTML_TEMPLATE.format(
        sync_time=sync_time,
        total_resources=total_resources,
        completed_steps=completed_steps,
        total_steps=total_steps,
        open_prs=open_prs,
        fully_reconciled=fully_reconciled,
        cards_html=cards_html
    )

    # Write to local repo location
    with open(OUTPUT_REPO_PATH, "w") as f:
        f.write(html_out)
    print(f"Generated dashboard inside repository: {OUTPUT_REPO_PATH}")

    # Write/Copy to App Data html_artifacts folder
    os.makedirs(os.path.dirname(OUTPUT_ARTIFACT_PATH), exist_ok=True)
    with open(OUTPUT_ARTIFACT_PATH, "w") as f:
        f.write(html_out)
    print(f"Generated dashboard inside App Data: {OUTPUT_ARTIFACT_PATH}")

def run_server(host, port):
    class DashboardHandler(BaseHTTPRequestHandler):
        def do_GET(self):
            if self.path in ["/", "/index.html"]:
                try:
                    with open(OUTPUT_REPO_PATH, "rb") as f:
                        content = f.read()
                    self.send_response(200)
                    self.send_header("Content-Type", "text/html")
                    self.send_header("Content-Length", str(len(content)))
                    self.end_headers()
                    self.wfile.write(content)
                except Exception as e:
                    self.send_error(500, f"Error reading dashboard file: {e}")
            else:
                self.send_error(404, "File not found")

        def do_POST(self):
            if self.path == "/sync":
                try:
                    print("Received sync request via HTTP...")
                    sync_and_generate()
                    self.send_response(200)
                    self.send_header("Content-Type", "application/json")
                    self.end_headers()
                    self.wfile.write(b'{"status": "success"}')
                except Exception as e:
                    print(f"Error during HTTP sync: {e}")
                    self.send_response(500)
                    self.send_header("Content-Type", "application/json")
                    self.end_headers()
                    self.wfile.write(json.dumps({"status": "error", "message": str(e)}).encode())
            else:
                self.send_error(404, "Not found")

    server_address = (host, port)
    httpd = HTTPServer(server_address, DashboardHandler)
    print(f"Serving dashboard at http://{host}:{port}/")
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        print("\nStopping server...")
        httpd.server_close()

def main():
    parser = argparse.ArgumentParser(description="Generate or serve the KCC Direct Migration Dashboard.")
    parser.add_argument("--serve", "-s", action="store_true", help="Start an HTTP server to serve the dashboard")
    parser.add_argument("--host", "-H", default="0.0.0.0", help="Host address to bind the server to (default: 0.0.0.0)")
    parser.add_argument("--port", "-p", type=int, default=8080, help="Port to bind the server to (default: 8080)")
    parser.add_argument("--no-sync", action="store_true", help="Skip initial sync with GitHub when generating or starting the server")
    args = parser.parse_args()

    if not args.no_sync:
        print("Running sync with GitHub...")
        sync_and_generate()
    else:
        print("Skipping GitHub sync. Generating dashboard from local database cache...")
        generate_dashboard_html()

    if args.serve:
        run_server(args.host, args.port)

if __name__ == "__main__":
    main()
