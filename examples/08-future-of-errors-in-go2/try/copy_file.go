package example

// Несуществующий синтаксис.

/*
func CopyFile(src, dst string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("copy %q to %q: %v", src, dst, err)
		}
	}()

	r := try(os.Open(src))
	defer r.Close()

	w := try(os.Create(dst))
	defer func() {
		_ = w.Close()

		if err != nil {
			_ = os.Remove(dst)
		}
	}()

	try(io.Copy(w, r))
	try(w.Close())
	return nil
}
*/
