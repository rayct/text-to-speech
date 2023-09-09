// A simple Text-to-Speech app
using System;
using System.Speech.Synthesis;

class Program
{
    static void Main()
    {
        // Create a SpeechSynthesizer instance
        using (SpeechSynthesizer synth = new SpeechSynthesizer())
        {
            // Text to be converted to speech
            string text = "Hello, this is a simple text-to-speech example in C#.";

            try
            {
                // Speak the text
                synth.Speak(text);
            }
            catch (Exception ex)
            {
                Console.WriteLine($"An error occurred: {ex.Message}");
            }
        }
    }
}

// using System;
// using Windows.Media.SpeechSynthesis;

// class Program
// {
//     static void Main(string[] args)
//     {
//         var synthesizer = new SpeechSynthesizer();
//         var text = "Hello, this is a test.";

//         // Generate the audio from text
//         var stream = synthesizer.SynthesizeTextToStream(text);

//         // Play the audio (you would need to handle this in your application)
//         // For example, you can save the stream to a file or use an audio player library.
//     }
// }


// Project and Documentation By: **Raymond C. TURNER**

// Last Updated: Saturday 9th September 2023