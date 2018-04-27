package data

import (
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/docker/docker/pkg/homedir"
	"log"
	"path/filepath"
	//"github.com/HouzuoGuo/tiedot/dberr"
)

func Database() *db.DB {
	dbPath := filepath.Join(homedir.Get(), ".eva", "data")

	// Open database (creates as necessary /w directories)
	myDB, err := db.OpenDB(dbPath)
	if err != nil {
		panic(err)
	}

	/* Create all needed collections.
	   ignore errors as we want to idompotently
	   create these if they do not exist. */
	myDB.Create("events")
	myDB.Create("responses")
	myDB.Create("invocations")

	return myDB
}

func PutEvent(event map[string]interface{}) int {
	myDB := Database()
	defer myDB.Close()
	events := myDB.Use("events")

	docID, err := events.Insert(event)
	if err != nil {
		log.Fatal(err)
	}
	return docID
}

func GetEvent(docID int) map[string]interface{} {
	myDB := Database()
	defer myDB.Close()
	events := myDB.Use("events")

	readBack, err := events.Read(docID)
	if err != nil {
		log.Fatal(err)
	}
	return readBack
}

/*
	// Update document
	err = feeds.Update(docID, map[string]interface{}{
		"name": "Go is very popular",
		"url":  "google.com"})
	if err != nil {
		panic(err)
	}

	// Process all documents (note that document order is undetermined)
	feeds.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		fmt.Println("Document", id, "is", string(docContent))
		return true  // move on to the next document OR
		return false // do not move on to the next document
	})

	// Delete document
	if err := feeds.Delete(docID); err != nil {
		panic(err)
	}

	// More complicated error handing - identify the error Type.
	// In this example, the error code tells that the document no longer exists.
	if err := feeds.Delete(docID); dberr.Type(err) == dberr.ErrorNoDoc {
		fmt.Println("The document was already deleted")
	}

	// ****************** Index Management ******************
	// Indexes assist in many types of queries
	// Create index (path leads to document JSON attribute)
	if err := feeds.Index([]string{"author", "name", "first_name"}); err != nil {
		panic(err)
	}
	if err := feeds.Index([]string{"Title"}); err != nil {
		panic(err)
	}
	if err := feeds.Index([]string{"Source"}); err != nil {
		panic(err)
	}

	// What indexes do I have on collection A?
	for _, path := range feeds.AllIndexes() {
		fmt.Printf("I have an index on path %v\n", path)
	}

	// Remove index
	if err := feeds.Unindex([]string{"author", "name", "first_name"}); err != nil {
		panic(err)
	}

	// ****************** Queries ******************
	// Prepare some documents for the query
	feeds.Insert(map[string]interface{}{"Title": "New Go release", "Source": "golang.org", "Age": 3})
	feeds.Insert(map[string]interface{}{"Title": "Kitkat is here", "Source": "google.com", "Age": 2})
	feeds.Insert(map[string]interface{}{"Title": "Good Slackware", "Source": "slackware.com", "Age": 1})

	var query interface{}
	json.Unmarshal([]byte(`[{"eq": "New Go release", "in": ["Title"]}, {"eq": "slackware.com", "in": ["Source"]}]`), &query)

	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, feeds, &queryResult); err != nil {
		panic(err)
	}

	// Query result are document IDs
	for id := range queryResult {
		// To get query result document, simply read it
		readBack, err := feeds.Read(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Query returned document %v\n", readBack)
	}

	// Gracefully close database
	if err := myDB.Close(); err != nil {
		panic(err)
	}
}
*/