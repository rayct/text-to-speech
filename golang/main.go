package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Text to be converted to speech
	text := "Your sentence."

	// Define file names for both WAV and MP3 files
	wavFileName := "output.wav"
	mp3FileName := "output.mp3"

	// Run the command to synthesize speech using htgo-tts and save it to a WAV file
	wavCmd := exec.Command("bash", "-c", fmt.Sprintf("htgo-tts -o %s -t \"%s\"", wavFileName, text))
	wavCmd.Stdout = os.Stdout
	wavCmd.Stderr = os.Stderr
	if err := wavCmd.Run(); err != nil {
		fmt.Printf("Failed to synthesize speech: %v\n", err)
		return
	}

	// Convert the WAV file to MP3 using an external tool (e.g., ffmpeg)
	mp3Cmd := exec.Command("ffmpeg", "-i", wavFileName, mp3FileName)
	mp3Cmd.Stdout = os.Stdout
	mp3Cmd.Stderr = os.Stderr
	if err := mp3Cmd.Run(); err != nil {
		fmt.Printf("Failed to convert WAV to MP3: %v\n", err)
		return
	}

	fmt.Printf("Speech saved as '%s' and '%s'\n", wavFileName, mp3FileName)
}

//////////////////////////////////

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	// Text to be converted to speech
// 	text := "Your sentence."

// 	// Define file names for both WAV and MP3 files
// 	wavFileName := "output.wav"
// 	mp3FileName := "output.mp3"

// 	// Run the command to synthesize speech using htgo-tts and save it to a WAV file
// 	wavCmd := exec.Command("./main", "-o", wavFileName, "-t", text)
// 	wavCmd.Stdout = os.Stdout
// 	wavCmd.Stderr = os.Stderr
// 	if err := wavCmd.Run(); err != nil {
// 		fmt.Printf("Failed to synthesize speech: %v\n", err)
// 		return
// 	}

// 	// Convert the WAV file to MP3 using an external tool (e.g., ffmpeg)
// 	mp3Cmd := exec.Command("ffmpeg", "-i", wavFileName, mp3FileName)
// 	mp3Cmd.Stdout = os.Stdout
// 	mp3Cmd.Stderr = os.Stderr
// 	if err := mp3Cmd.Run(); err != nil {
// 		fmt.Printf("Failed to convert WAV to MP3: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Speech saved as '%s' and '%s'\n", wavFileName, mp3FileName)
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	// Text to be converted to speech
// 	text := "Your sentence."

// 	// Run the command to synthesize speech using htgo-tts and save it to a WAV file
// 	wavFileName := "output.wav"
// 	cmd := exec.Command("bash", "-c", fmt.Sprintf("htgo-tts -o %s -t \"%s\"", wavFileName, text))
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Run(); err != nil {
// 		fmt.Printf("Failed to synthesize speech: %v\n", err)
// 		return
// 	}

// 	// Convert the WAV file to MP3 using an external tool (e.g., ffmpeg)
// 	mp3FileName := "output.mp3"
// 	cmd = exec.Command("ffmpeg", "-i", wavFileName, mp3FileName)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Run(); err != nil {
// 		fmt.Printf("Failed to convert WAV to MP3: %v\n", err)
// 		return
// 	}

// 	// Remove the temporary WAV file
// 	if err := os.Remove(wavFileName); err != nil {
// 		fmt.Printf("Failed to remove temporary WAV file: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Speech saved as '%s'\n", mp3FileName)
// }

// Project and Documentation by: **Raymond C. TURNER**

// Last Updated: Saturday 9th September 2023
