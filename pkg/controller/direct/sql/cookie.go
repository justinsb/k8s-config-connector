package sql

// // updateCookie computes a hash of the instance and stores it in an annotation, so that we can detect drift in the next reconciliation loop.
// func updateCookie(ctx context.Context, instance *api.DatabaseInstance, updateOp *directbase.UpdateOperation) error {
// 	h := hasher{inner: fnv.New128a()}

// 	h.appendJSON(instance)
// 	instanceHash := h.HashString()

// 	h.appendJSON(spec)
// 	specHash := h.HashString()

// 	hash := instanceHash + ":" + specHash

// 	if h.Err() != nil {
// 		return h.Err()
// 	}

// 	updateOp.SetAnnotation(k8s.CookieAnnotation, hash)
// 	return nil
// }

// type hasher struct {
// 	inner  hash.Hash
// 	errors []error
// }

// // Reset resets the hasher to its initial state. It does not clear the errors, as they are only relevant to the current hashing operation.
// func (h hasher) Reset() {
// 	h.inner.Reset()
// }

// func (h hasher) Err() error {
// 	if len(h.errors) == 0 {
// 		return nil
// 	}
// 	return errors.Join(h.errors...)
// }

// func (h hasher) HashString() string {
// 	hashBytes := h.inner.Sum(nil)
// 	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashBytes)
// }

// func (h hasher) appendProto(obj proto.Message) {
// 	b, err := proto.MarshalOptions{Deterministic: true}.Marshal(obj)
// 	if err != nil {
// 		h.errors = append(h.errors, err)
// 		return
// 	}

// 	h.appendWithLength(b)
// }

// func (h hasher) appendJSON(obj any) {
// 	b, err := json.Marshal(obj)
// 	if err != nil {
// 		h.errors = append(h.errors, err)
// 		return
// 	}

// 	h.appendWithLength(b)
// }

// // appendWithLength
// func (h hasher) appendWithLength(b []byte) {
// 	lenBuffer := make([]byte, 8)
// 	binary.BigEndian.PutUint64(lenBuffer, uint64(len(b)))

// 	if _, err := h.inner.Write(lenBuffer); err != nil {
// 		h.errors = append(h.errors, err)
// 		return
// 	}
// 	if _, err := h.inner.Write(b); err != nil {
// 		h.errors = append(h.errors, err)
// 		return
// 	}
// }
