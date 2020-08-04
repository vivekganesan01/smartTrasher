package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
)

type configuration struct {
	DownloadDir  string `yaml:"DownloadDirectory"`
	AutoTrashDir string `yaml:"AutoTrashDirectory"`
	Day          int32  `yaml:"RecyclingDays"`
}

var expireTime float64
var today time.Time

var dir [5]string = [5]string{"images", "documents", "zip", "dir", "miscellaneous"}
var img = [...]string{".png", ".jpeg", ".jpg", ".svn"}
var doc = [...]string{".doc", ".xlsx", ".xls", ".pdf", ".log", ".csv", ".log", ".txt", ".jmx", ".docx", ".html", ".xml", ".json"}
var zip = [...]string{".zip", ".dmg", ".gz"}

// Helps to verify whether the file is modified within expireTime frame.
func timeMachine(lastModified time.Time) bool {
	if today.Sub(lastModified).Hours() > expireTime {
		return true
	}
	return false
}

// moveTo helps to move to file from download dir to custom dir with AutoTrash dir
func (c *configuration) moveTo(fileName string, dir string) {
	fmt.Println("Last modified "+strconv.Itoa(int(c.Day))+" days ago, compressing to:"+c.DownloadDir+"/"+fileName, c.AutoTrashDir+"/"+dir+"/"+fileName)
	os.Rename(c.DownloadDir+"/"+fileName, c.AutoTrashDir+"/"+dir+"/"+fileName)
}

// run helps to filter out the files based on spec provided by *.yml
func (c *configuration) run() {
	expireTime = float64(c.Day) * 24 // number of days before the current day
	today = time.Now()               // Store the current exec time stamp

	// To make sure all trash dir exists ("images", "documents", "zip", "dir", "miscellaneous")
	for i := 0; i < len(dir); i++ {
		err := os.MkdirAll(c.AutoTrashDir+"/"+dir[i], 0755)
		if err != nil {
			fmt.Println("Failed @ dir creation", err.Error())
		}
		fmt.Println("Created Dir:", dir[i])
	}

	// loops through download directory
	fileINfo, readErr := ioutil.ReadDir(c.DownloadDir)
	if readErr != nil {
		panic(fmt.Sprintf("Unable to read dir: %s, %s", c.DownloadDir, readErr.Error()))
	}

	// validates each file's modification time against current exec time
	for _, entry := range fileINfo {
		if !entry.IsDir() {
			fileName := entry.Name()
			// documents
			for _, idoc := range doc {
				if strings.Contains(fileName, idoc) {
					if timeMachine(entry.ModTime()) {
						c.moveTo(fileName, "documents")
					}
					break
				}
			}
			// images
			for _, image := range img {
				if strings.Contains(fileName, image) {
					if timeMachine(entry.ModTime()) {
						c.moveTo(fileName, "images")
					}
					break
				}
			}
			// zip
			for _, izip := range zip {
				if strings.Contains(fileName, izip) {
					if timeMachine(entry.ModTime()) {
						c.moveTo(fileName, "zip")
					}
					break
				}
			}
			// miscellaneous
			if timeMachine(entry.ModTime()) {
				c.moveTo(fileName, "miscellaneous")
			}
		}
	}
	fmt.Println(" - * End * -")
}

// MAIN handler
func main() {
	start := time.Now()
	fmt.Println("Welcome to Auto Trash Algo ....")
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic("please make sure config.yml file is in place .... Exiting 1")
	}
	config := configuration{}
	yaml.Unmarshal(yamlFile, &config)
	config.run()
	duration := time.Since(start)
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Println(" - * -------------------- * -")
}
