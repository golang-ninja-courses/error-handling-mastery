package example

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFileV3(src, dst string) (err error) {
	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("copy %q to %q: %v", src, dst, err)
	}
	defer func() {
		if err2 := r.Close(); err2 != nil {
			log.Printf("copy %q to %q: cannot close src file: %v", src, dst, err2)
		}
	}()

	w, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copy %q to %q: %v", src, dst, err)
	}
	defer func() {
		if wErr := w.Close(); wErr != nil {
			log.Printf("copy %q to %q: cannot close dst file: %v", src, dst, wErr)
		}

		if err != nil {
			if err2 := os.Remove(dst); err2 != nil {
				log.Printf("copy %q to %q: cannot remove dst file : %v", src, dst, err2)
			}
		}
	}()

	if _, err := io.Copy(w, r); err != nil {
		return fmt.Errorf("copy %q to %q: cannot do io.Copy: %v", src, dst, err)
	}

	if err := w.Sync(); err != nil {
		return fmt.Errorf("copy %q to %q: cannot sync dst file %v", src, dst, err)
	}

	return nil
}
