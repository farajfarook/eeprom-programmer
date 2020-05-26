#include <Arduino.h>
#include <SoftwareSerial.h>

#define PIN_RX PB3
#define PIN_TX PB4
#define PIN_LED PB0

#define CMD_ON "on"
#define CMD_OFF "off"

SoftwareSerial mySerial(PIN_RX, PIN_TX);

void setup()
{
    mySerial.begin(19200);
    pinMode(PIN_LED, OUTPUT);
}

void loop()
{
    char buffer[16];
    while (!mySerial.available())
        ;

    int size = mySerial.readBytesUntil('\n', buffer, 16);
    if (size == 0)
        return;

    if (strncmp(buffer, CMD_ON, strlen(CMD_ON)) == 0)
    {
        mySerial.println(buffer);
        digitalWrite(PIN_LED, HIGH);
    }
    else if (strncmp(buffer, CMD_OFF, strlen(CMD_OFF)) == 0)
    {
        mySerial.println(buffer);
        digitalWrite(PIN_LED, LOW);
    }
    else
    {
        char output[50];
        sprintf(output, "unknown command: %s", buffer);
        mySerial.println(output);
    }
}