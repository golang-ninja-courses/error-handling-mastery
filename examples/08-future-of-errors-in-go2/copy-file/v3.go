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
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copy %q to %q: create dst file: %v", src, dst, err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			log.Printf("copy %q to %q: cannot close dst file: %v", src, dst, err)
		}

		if err != nil {
			if err := os.Remove(dst); err != nil {
				log.Printf("copy %q to %q: cannot remove dst file : %v", src, dst, err)
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
