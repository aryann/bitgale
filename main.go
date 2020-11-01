package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aryann/bencode"
)

// MetaInfo is the data schema for .torrent files. For more information, refer
// to https://wiki.theory.org/BitTorrentSpecification.
type MetaInfo struct {
	Info         Info       `bencode:"info"`
	Announce     string     `bencode:"announce"`
	AnnounceList [][]string `bencode:"announce-list"`
	CreationDate int64      `bencode:"creation date"`
	Comment      string     `bencode:"comment"`
	CreatedBy    string     `bencode:"created by"`
	Encoding     string     `bencode:"encoding"`
}

type Info struct {
	Name        string `bencode:"name"`
	PieceLength int64  `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Private     int64  `bencode:"private"`

	// Single-file mode fields.
	Length string `bencode:"length"`
	MD5Sum string `bencode:"md5sum"`

	// Multi-file mode fields.
	Files []File `bencode:"files"`
}

type File struct {
	Length int64    `bencode:"length"`
	MD5Sum string   `bencode:"md5sum"`
	Path   []string `bencode:"path"`
}

func main() {
	flag.Parse()

	filePath := flag.Arg(0)
	if filePath == "" {
		log.Fatalf("A .torrent file path must be provided")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Could not read from file: %v", err)
	}

	metaInfo := MetaInfo{}
	if err := bencode.Unmarshal(data, &metaInfo); err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}

	output, err := json.MarshalIndent(metaInfo, "", "  ")
	if err != nil {
		log.Fatalf("Could not dump to JSON: %v", err)
	}
	fmt.Println(string(output))
}
