package logs

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type MuxWriter struct {
	sync.Mutex
	fd *os.File
}

func (this *MuxWriter) Write(b []byte) (int, error) {
	this.Lock()
	defer this.Unlock()
	return this.fd.Write(b)
}

func (this *MuxWriter) SetFd(fd *os.File) {
	if this.fd != nil {
		this.fd.Close()
	}
	this.fd = fd
}

type FileLogWriter struct {
	*log.Logger
	mw *MuxWriter

	Filename    string `json:"filename"`
	Level       int    `json:"level"`
	curFilename string
	daily       int

	startLock sync.Mutex // Only one log can write to the file
}

func NewFileWriter() LoggerInterface {
	fw := &FileLogWriter{
		Filename: "",
		Level:    LevelDebug,
	}
	fw.mw = new(MuxWriter)
	fw.Logger = log.New(fw.mw, "", log.Ldate|log.Ltime)
	return fw
}

func (this *FileLogWriter) Init(config string) error {
	err := json.Unmarshal([]byte(config), this)
	if err != nil {
		return err
	}
	if len(this.Filename) == 0 {
		return errors.New("jsonconfig must have filename")
	}
	err = this.startLogger()
	return err
}

func (this *FileLogWriter) startLogger() error {
	fd, err := this.createLogFile()
	if err != nil {
		return err
	}
	this.mw.SetFd(fd)
	this.daily = time.Now().Day()
	return nil
}

func (this *FileLogWriter) createLogFile() (*os.File, error) {
	this.curFilename = this.Filename + fmt.Sprintf(".%v", time.Now().Format("2006-01-02"))
	fd, err := os.OpenFile(this.curFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0660)
	return fd, err
}

func (this *FileLogWriter) WriteMsg(msg string, level int) error {
	if level > this.Level {
		return nil
	}
	n := 24 + len(msg) // 24 stand for the length "2013/06/23 21:00:22 [T] "
	this.docheck(n)
	this.Logger.Println(msg)
	return nil
}

func (this *FileLogWriter) docheck(size int) {
	this.startLock.Lock()
	defer this.startLock.Unlock()
	if time.Now().Day() != this.daily {
		if err := this.DoRotate(); err != nil {
			fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", this.curFilename, err)
			return
		}
	}
}

func (this *FileLogWriter) DoRotate() error {
	this.mw.Lock()
	defer this.mw.Unlock()

	this.Destroy()

	err := this.startLogger()
	if err != nil {
		return fmt.Errorf("Rotate StartLogger: %s\n", err)
	}

	return nil
}

func (this *FileLogWriter) Destroy() {
	this.mw.fd.Close()
}

func (this *FileLogWriter) Flush() {
	this.mw.fd.Sync()
}

func init() {
	register("file", NewFileWriter)
}
