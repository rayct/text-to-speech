#include <stdio.h>
#include <espeak/speak_lib.h>

int main()
{
    if (espeak_Initialize(AUDIO_OUTPUT_PLAYBACK, 0, NULL, 0) < 0)
    {
        printf("Failed to initialize eSpeak.\n");
        return 1;
    }

    const char *text = "Hello, this is a simple text-to-speech example in C.";
    espeak_Synth(text, strlen(text) + 1, 0, POS_CHARACTER, 0, espeakCHARS_AUTO, NULL, NULL);
    espeak_Synchronize();

    espeak_Terminate();

    return 0;
}

// Project and Documentation by: **Raymond C. TURNER**

// Last Updated: Saturday 9th September 2023