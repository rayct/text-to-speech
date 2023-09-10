from gtts import gTTS
import os

# Text to be converted to speech
text = "Return Values and Input Parameters of Methods in C#. In C#, methods can have return values and input parameters that allow them to interact with other parts of your code. Some methods are designed to perform a specific task and end quietly without returning a value. These methods are referred to as 'void methods.' On the other hand, some methods are designed to compute a result and return that value upon completion. This return value typically represents the outcome of an operation and serves as the primary means for a method to communicate back to the code that calls it."

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
# os.remove("output.mp3")


# Project and Documentation by: **Raymond C. TURNER**

# Last Updated: Monday 11th September 2023
