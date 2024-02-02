package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type RotateHook struct {
	FileName string
	OpenDate int
	MaxDays  int64
}

func NewRotateHook(fileName string, maxDays int64) *RotateHook {
	return &RotateHook{FileName: fileName, MaxDays: maxDays, OpenDate: time.Now().Day()}
}

func (r *RotateHook) Fire(w *io.PipeWriter) (f *os.File, err error) {
	if time.Now().Day() == r.OpenDate {
		return nil, nil
	}

	num := 1
	fileName := ""
	// find the next available number
	for ; err == nil && num <= 999; num++ {
		fileName = r.FileName + fmt.Sprintf(".%s.%03d", time.Now().Format(rotateFileNameLayout), num)
		_, err = os.Lstat(fileName)
	}

	_, err = os.Lstat(r.FileName)
	if err != nil {
		return nil, err
	}

	// closed fd before rename
	err = w.Close()
	if err != nil {
		return nil, err
	}

	// rename the file to it's newfound home
	err = os.Rename(r.FileName, fileName)
	if err != nil {
		return nil, fmt.Errorf("rotate error: %w\n", err)
	}

	// reopen file
	f, err = os.OpenFile(r.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	r.OpenDate = time.Now().Day()

	go func() {
		_ = filepath.Walk(filepath.Dir(r.FileName), r.removeExpired)
	}()

	return f, nil
}

// removeExpired returns a filepath.WalkFunc for filepath.Walk.
func (r *RotateHook) removeExpired(path string, info os.FileInfo, err error) error {
	var walkErr error
	defer func() {
		if val := recover(); val != nil {
			walkErr = fmt.Errorf("unable to del old log %q, err: %+v", path, val)
		}

		if walkErr != nil {
			fmt.Println(walkErr)
		}
	}()

	if !info.IsDir() && info.ModTime().Unix() < (time.Now().Unix()-60*60*24*r.MaxDays) {
		if strings.HasPrefix(filepath.Base(path), filepath.Base(r.FileName)) {
			walkErr = os.Remove(path)
		}
	}

	return walkErr
}
