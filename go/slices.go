package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

// 2 ways of doing it
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			// matrix[i] = append(matrix[i], i*j)
			matrix[i][j] = i * j
		}
	}

	return matrix
}

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	var filtered []Message
	for _, message := range messages {
		if message.Type() == filterType {
			filtered = append(filtered, message)
		}
	}
	return filtered
}

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// fmt.Println("======================")
	for _, msg := range messages {
		msg.tags = make([]string, 1)
		msg.tags = tagger(msg)
		// fmt.Println(i, msg.tags)
	}
	// fmt.Println("======================")
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	// fmt.Println("======================")
	for _, c := range msg.content {
		fmt.Println(c)
	}
	// fmt.Println("======================")
	return tags
}
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(first, second string) {
		fmt.Println(formatter(first, second))
	}
}

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}
