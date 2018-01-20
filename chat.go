package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	logSeparator = "\xbd"
)

type Chat struct {
	M string
	T time.Time
}

func (c *Chat) String() string {
	return fmt.Sprintf("%v"+logSeparator+"%v", c.T.String(), c.M)
}

func NewChat(m string) *Chat {
	return &Chat{m, time.Now()}
}

func GetLog(filename string, t time.Time) (error, Chat) {
	return errors.New("Bad"), Chat{}
}

func logToMap(filename string) (map[time.Time]string, error) {
	err := logFileValid(filename)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fileLen, err := f.Seek(0, 2)
	if err != nil {
		return nil, err
	}
	fileBytes := make([]byte, fileLen)
	f.Read(fileBytes)
	logLines := bytes.Split(fileBytes, []byte(logSeparator))
	return make(map[time.Time]string, 1), nil
}

func (c *Chat) LogToFile(filename string) error {
	if err := logFileValid(filename); err == nil {
		// time.Date(year, month, day, hour, min, sec, nsec, loc)
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		defer f.Close()
		if err != nil {
			return err
		}
		n, err := f.WriteString(c.String() + "\n")
		if err != nil {
			return err
		}
		err = f.Sync()
		if err != nil {
			return err
		}
		log.Printf("%v bytes written to log file %v", n, filename)
	} else {
		return err
	}
	return nil
}

func logFileValid(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("File %s does not exist", filename))
	} else {
		return nil
	}
}
