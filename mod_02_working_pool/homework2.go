package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

func main() {
	rand.Seed(time.Now().Unix())

	const usersCount, workerCount = 1000, 1

	startTime := time.Now()

	usersIndices := make(chan int, usersCount)
	users := make(chan User, usersCount)
	finish := make(chan bool, usersCount)

	for i := 0; i < workerCount; i++ {
		go generateUsers(usersIndices, users)
	}

	for i := 0; i < workerCount; i++ {
		go saveUsersInfo(users, finish)
	}

	for i := 0; i < usersCount; i++ {
		usersIndices <- i
	}

	for i := 0; i < usersCount; i++ {
		<-finish
	}

	close(usersIndices)
	close(users)
	close(finish)

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUsersInfo(users <-chan User, finish chan<- bool) {
	for user := range users {
		saveUserInfo(user)
		finish <- true
	}
}

func saveUserInfo(user User) {
	fmt.Printf("WRITING FILE FOR UID %d\n", user.id)

	filename := fmt.Sprintf("users/uid%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(user.getActivityInfo())
	time.Sleep(time.Second * 1)
}

func generateUser(index int) User {
	user := User{
		id:    index + 1,
		email: fmt.Sprintf("user%d@company.com", index+1),
		logs:  generateLogs(rand.Intn(1000)),
	}
	fmt.Printf("generated user %d\n", index+1)
	time.Sleep(time.Millisecond * 100)

	return user
}

func generateUsers(usersIndices <-chan int, users chan<- User) {
	for i := range usersIndices {
		users <- generateUser(i)
	}
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
