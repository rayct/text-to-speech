# Text-to-Speech (TTS)
## Python

A simple example of a text-to-speech program in Python using the `gTTS` (Google Text-to-Speech) library, which allows you to convert text into speech using Google Translate's TTS engine. You'll need to install the `gTTS` library before running the code:

This Python example provides a basic text-to-speech functionality using the `gTTS` library, which is simple to use for quick TTS tasks. If you need more advanced TTS capabilities, you may consider using other Python libraries or cloud-based TTS services like the Google Cloud Text-to-Speech API.

```bash
pip install gTTS
```

Here's the Python code:

```python
from gtts import gTTS
import os

# Text to be converted to speech
text = "Hello, this is a simple text-to-speech example in Python."

# Create a gTTS object
tts = gTTS(text)

# Save the generated speech to a file (optional)
tts.save("output.mp3")

# Play the generated speech (optional)
os.system("mpg321 output.mp3")  # You may need to install mpg321 if not already installed

# Alternatively, you can play the speech directly without saving it
# tts.save("output.mp3")
# os.system("mpg321 output.mp3")

# Clean up the temporary speech file (optional)
os.remove("output.mp3")
```

In this Python code:

1. We import the necessary libraries, including `gTTS` for text-to-speech conversion.
2. We provide the text message to be converted to speech in the `text` variable.
3. We create a `gTTS` object with the text.
4. We can optionally save the generated speech to an MP3 file using `tts.save("output.mp3")`.
5. We can optionally play the generated speech using the `os.system()` function. Make sure you have a command-line MP3 player like `mpg321` installed for this to work.
6. Optionally, you can clean up the temporary speech file if you choose to save it.

Run the Python script:

```bash
python tts_example.py
```

</br>

---

Project Developer and Documentation By: **Raymond C. TURNER**

Last Updated: Saturday 9th September 2023