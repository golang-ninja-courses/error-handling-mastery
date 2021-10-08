package example

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(src, dst string) (err error) {
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

// Несуществующий синтаксис.

/*
func CopyFile(src, dst string) (err error) {
	handle err {
		return fmt.Errorf("copy %q to %q: %v", src, dst, err)
	}

	r := check os.Open(src)
	defer r.Close()

	handle err {
		return fmt.Errorf("copy %q to %q: create dst file: %v", src, dst, err)
	}
	w := check os.Create(dst)
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

	handle err {
		return fmt.Errorf("copy %q to %q: cannot do io.Copy: %v", src, dst, err)
	}
	check io.Copy(w, r)

	handle err {
		return fmt.Errorf("copy %q to %q: cannot sync dst file %v", src, dst, err)
	}
	check w.Sync()

	return nil
}
*/

// Несуществующий синтаксис.

/*
func CopyFile(src, dst string) (err error) {
	handle err {
		return fmt.Errorf("copy %q to %q: %v", src, dst, err)
	}

	r := check os.Open(src)
	defer r.Close()

	w := check os.Create(dst)
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

	check io.Copy(w, r)
	check w.Sync()
	return nil
}
*/

// Несуществующий синтаксис.

/*
func CopyFile(src, dst string) (err error) {
	handle err {
		return fmt.Errorf("copy %q to %q: %v", src, dst, err)
	}

	r := check os.Open(src)
	defer r.Close()

	w := check os.Create(dst)
	defer func() {
        _ = w.Close()

		if err != nil {
			_ = os.Remove(dst)
		}
	}()

	check io.Copy(w, r)
	check w.Sync()
	return nil
}
*/
