package utils

import (
	"encoding/gob"
	"fmt"
	"math"
	"os"
	"strings"
)

// IntentExample represents a training example with text and its corresponding intent.
type IntentExample struct {
	Text   string
	Intent string
}
type ClassificationResponse struct {
	Decision string `json:"decision"`
}

// NaiveBayesClassifier is the structure to hold the training results.
type NaiveBayesClassifier struct {
	WordProbabilities   map[string]map[string]float64
	IntentProbabilities map[string]float64
	Vocabulary          map[string]bool
}

// TrainNaiveBayes trains a Naive Bayes classifier given training data and returns the model.
func TrainNaiveBayes(data []IntentExample) NaiveBayesClassifier {
	wordCounts := make(map[string]map[string]int)
	intentCounts := make(map[string]int)
	vocabulary := make(map[string]bool)
	totalExamples := len(data)

	// Initialize data structures
	for _, example := range data {
		intent := example.Intent
		if _, exists := wordCounts[intent]; !exists {
			wordCounts[intent] = make(map[string]int)
		}
		intentCounts[intent]++

		// Tokenize and normalize the text
		words := strings.Fields(strings.ToLower(example.Text))
		for _, word := range words {
			wordCounts[intent][word]++
			vocabulary[word] = true
		}
	}

	// Calculate probabilities
	wordProbabilities := make(map[string]map[string]float64)
	intentProbabilities := make(map[string]float64)
	for intent, words := range wordCounts {
		wordProbabilities[intent] = make(map[string]float64)
		sumWords := 0
		for _, count := range words {
			sumWords += count
		}
		for word := range vocabulary {
			prob := (float64(words[word]) + 1) / float64(sumWords+len(vocabulary)) // Laplace smoothing
			wordProbabilities[intent][word] = prob
		}
		intentProbabilities[intent] = float64(intentCounts[intent]) / float64(totalExamples)
	}

	return NaiveBayesClassifier{
		WordProbabilities:   wordProbabilities,
		IntentProbabilities: intentProbabilities,
		Vocabulary:          vocabulary,
	}
}

// Classify uses the trained Naive Bayes classifier to determine the most likely intent for a given text.
func (classifier *NaiveBayesClassifier) Classify(text string) ClassificationResponse {
	maxProbability := math.Inf(-1)
	selectedIntent := "unknown"
	words := strings.Fields(strings.ToLower(text))

	for intent, intentProb := range classifier.IntentProbabilities {
		prob := math.Log(intentProb)
		for _, word := range words {
			if _, exists := classifier.Vocabulary[word]; exists {
				prob += math.Log(classifier.WordProbabilities[intent][word])
			} else {
				prob += math.Log(1 / float64(len(classifier.Vocabulary)+1)) // Handle unknown words
			}
		}
		if prob > maxProbability {
			maxProbability = prob
			selectedIntent = intent
		}
	}

	return ClassificationResponse{
		Decision: selectedIntent,
	}
}

// SaveModel saves the NaiveBayesClassifier to a file using gob encoding.
func SaveModel(classifier NaiveBayesClassifier, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(classifier); err != nil {
		return fmt.Errorf("error encoding model: %v", err)
	}
	return nil
}

// LoadModel loads the NaiveBayesClassifier from a file using gob encoding.
func LoadModel(filePath string) (NaiveBayesClassifier, error) {
	var classifier NaiveBayesClassifier
	file, err := os.Open(filePath)
	if err != nil {
		return classifier, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&classifier); err != nil {
		return classifier, fmt.Errorf("error decoding model: %v", err)
	}
	return classifier, nil
}

func TrainModel() error {
	trainingData := []IntentExample{
		{Text: "Your SORA application has been reviewed and approved. Congratulations on meeting all the necessary safety standards.", Intent: "Approval"},
		{Text: "After thorough examination, we are pleased to announce that your operation has been granted approval.", Intent: "Approval"},
		{Text: "All criteria have been successfully met. Your SORA application is approved.", Intent: "Approval"},
		{Text: "Please provide further details about the risk mitigation strategies mentioned in your SORA application.", Intent: "More information needed"},
		{Text: "We need additional documentation on your safety protocols before proceeding with the approval.", Intent: "More information needed"},
		{Text: "Your application is pending until further evidence of operational compliance is submitted.", Intent: "More information needed"},
		{Text: "Due to several unaddressed high-risk factors, we are unable to approve your SORA application.", Intent: "Declined"},
		{Text: "After careful review, we must decline your application as it does not meet the safety and operational standards required.", Intent: "Declined"},
		{Text: "Your application is declined due to insufficient risk mitigation measures.", Intent: "Declined"},
	}

	// Train the classifier
	classifier := TrainNaiveBayes(trainingData)
	// Save the trained model to a file
	modelFilePath := "trained_model.gob"
	if err := SaveModel(classifier, modelFilePath); err != nil {
		fmt.Println("Error saving model:", err)
		return err
	}
	return nil
}

func GetDecision(data string) (error, *string) {

	modelFilePath := "trained_model.gob"

	loadedClassifier, err := LoadModel(modelFilePath)
	if err != nil {
		fmt.Println("Error loading model:", err)
		return err, nil
	}

	response := loadedClassifier.Classify(data)

	return nil, &response.Decision

}
