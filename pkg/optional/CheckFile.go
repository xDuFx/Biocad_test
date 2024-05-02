package optional

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testB/pkg/models"
	"testB/pkg/repository"

	"github.com/fsnotify/fsnotify"
)

func ParseAndCheckTcv(fileDirect string, fileDirectPdf string, db *repository.PGRepo) {

	listDb, err := db.GetFilenames()
	if err != nil {
		log.Fatal(err)
	}
	listTsv, err := os.ReadDir(fileDirect)
	if err != nil {
		log.Fatal(err)
	}
	createCh := make(chan string, 10)
	for _, file := range listTsv {
		flag := false
		for _, fileDb := range listDb {
			if file.Name() == fileDb {
				flag = true
				break
			}
		}
		if !flag {
			if CheckResol(file.Name()) {
				fullname := fileDirect + "/" + file.Name()
				createCh <- fullname
			} else {
				RecordingError("Wrong resolution "+file.Name(), nil, db)
			}
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	err = watcher.Add(fileDirect)
	if err != nil {
		log.Fatalln(err)
	}
	dataCh := make(chan []models.DataTcv, 5)
	renameCh := make(chan bool)
	errCh := make(chan error)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				switch {
				case event.Has(fsnotify.Create):
					if !CheckResol(event.Name) {
						RecordingError("Wrong resolution "+event.Name, nil, db)
						continue
					}
					listDb, err := db.GetFilenames()
					if err != nil {
						log.Fatal(err)
					}
					flag := false
					for _, fileDb := range listDb {
						if fileDb == NameSplit(event.Name) {
							flag = true
							break
						}
					}
					if !flag {
						createCh <- event.Name
					}
				case event.Has(fsnotify.Rename):
					renameCh <- true
				}
			case err := <-watcher.Errors:
				errCh <- err
			}
		}
	}()

	go func() {
		for {
			select {
			case <-renameCh:

				err = watcher.Add(fileDirect)
				if err != nil {
					log.Fatalln(err)
				}
			case name := <-createCh:

				data := ParseTsv(name, db)
				db.NewData(*data)
				db.NewFilename(NameSplit(name))
				dataCh <- *data
				err = watcher.Add(fileDirect)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()

	go func() {
		for {
			data := <-dataCh
			err := CreateF(fileDirectPdf, data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	log.Fatalln(<-errCh)
}

func ParseTsv(fileDirect string, db *repository.PGRepo) *[]models.DataTcv {
	f, err := os.Open(fileDirect)
	if err != nil {
		RecordingError("Error open file "+fileDirect, err, db)
	}
	r := csv.NewReader(f)
	r.Comma = '\t'
	data := []models.DataTcv{}
	for count := 0; ; count++ {
		rows, err := r.Read()
		dataStr := []string{}
		if err == io.EOF {
			break
		}
		if count < 2 {
			continue
		}
		for i := range rows {
			row := rows[i]
			dataStr = append(dataStr, strings.Trim(row, "           "))
		}
		row := models.DataTcv{}
		row.N, err = strconv.Atoi(dataStr[0])
		if err != nil {
			RecordingError("Wrong format field N "+fileDirect, err, db)
			os.Exit(1)
		}
		row.Mqtt = dataStr[1]
		row.Invid = dataStr[2]
		row.Unit_guid = dataStr[3]
		row.Msg_id = dataStr[4]
		row.Text = dataStr[5]
		row.Context = dataStr[6]
		row.Class = dataStr[7]
		row.Level, err = strconv.Atoi(dataStr[8])
		if err != nil {
			RecordingError("Wrong format field Level"+fileDirect, err, db)
			os.Exit(1)
		}
		row.Area = dataStr[9]
		row.Addr = dataStr[10]
		row.Block = dataStr[11]
		row.TypE = dataStr[12]
		row.Bit = dataStr[13]
		row.Invert_bit = dataStr[14]
		data = append(data, row)
	}
	return &data
}
