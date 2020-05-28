/*
 * Blink
 * Turns on an LED on for one second,
 * then off for one second, repeatedly.
 */
#include "./pins.h"

void onClockTick();

void setup()
{
    Serial.begin(9600);
    for (size_t i = 0; i < DATA_LEN; i++)
    {
        pinMode(DATA[i], INPUT);
    }
    pinMode(CLK_INTR, INPUT);

    attachInterrupt(digitalPinToInterrupt(CLK_INTR), onClockTick, RISING);
}

void loop()
{
}

void onClockTick()
{
    unsigned int data = 0;
    int spaceCount = 0;
    for (size_t i = 0; i < DATA_LEN; i++)
    {
        int bit = digitalRead(DATA[i]) ? 1 : 0;
        Serial.print(bit);
        data = (data << 1) + bit;
        if (++spaceCount >= 4)
        {
            Serial.print(' ');
            spaceCount = 0;
        }
    }
    Serial.println();
}