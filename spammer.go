package goTp

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	var wg sync.WaitGroup

	in := make(chan interface{})

	for _, c := range cmds {
		out := make(chan interface{})

		wg.Add(1)

		go func(cmd cmd, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			cmd(in, out)
		}(c, in, out)

		in = out
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	seenUsers := make(map[uint64]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for email := range in {
		emailStr, ok := email.(string)
		if !ok {
			log.Printf("Invalid type for email: %T, value: %v", email, email)
			return
		}

		wg.Add(1)

		go func(email string) {
			defer wg.Done()

			user := GetUser(email)

			mu.Lock()
			defer mu.Unlock()

			if !seenUsers[user.ID] {
				seenUsers[user.ID] = true
				out <- user
			}
		}(emailStr)
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	usersBatch := []User{}

	for data := range in {
		user, ok := data.(User)
		if !ok {
			log.Printf("Invalid type for user: %T, value: %v", data, data)
			return
		}

		mu.Lock()
		usersBatch = append(usersBatch, user)
		if len(usersBatch) == GetMessagesMaxUsersBatch {
			batch := usersBatch
			usersBatch = nil
			mu.Unlock()

			wg.Add(1)
			go func(batch []User) {
				defer wg.Done()
				messages, err := GetMessages(batch...)
				if err != nil {
					log.Printf("Error getting messages for batch: %v", err)
					return
				}
				for _, msg := range messages {
					out <- msg
				}
			}(batch)
		} else {
			mu.Unlock()
		}
	}

	mu.Lock()
	if len(usersBatch) > 0 {
		wg.Add(1)
		go func(batch []User) {
			defer wg.Done()
			messages, err := GetMessages(batch...)
			if err != nil {
				log.Printf("Error getting messages for remaining batch: %v", err)
				return
			}
			for _, msg := range messages {
				out <- msg
			}
		}(usersBatch)
	}
	mu.Unlock()

	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, HasSpamMaxAsyncRequests)

	for data := range in {
		msgID, ok := data.(MsgID)
		if !ok {
			log.Printf("Invalid type for MsgID: %T, value: %v", data, data)
			return
		}

		wg.Add(1)
		go func(msgID MsgID) {
			defer wg.Done()

			sem <- struct{}{}
			hasSpam, err := HasSpam(msgID)
			if err != nil {
				log.Printf("Error checking spam for MsgID %v: %v", msgID, err)
				return
			}
			<-sem
			out <- MsgData{ID: msgID, HasSpam: hasSpam}
		}(msgID)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var results []MsgData

	for data := range in {
		msgData, ok := data.(MsgData)
		if ok {
			results = append(results, msgData)
		} else {
			log.Printf("Invalid type for MsgData: %T, value: %v", data, data)
			return
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].HasSpam == results[j].HasSpam {
			return results[i].ID < results[j].ID
		}
		return results[i].HasSpam
	})
	
	for _, result := range results {
		out <- fmt.Sprintf("%v %v", result.HasSpam, result.ID)
	}
}
