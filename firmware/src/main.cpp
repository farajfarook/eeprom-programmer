#include <Arduino.h>
#include <SoftwareSerial.h>

#define PIN_TX PB3
#define PIN_RX PB2

#define PIN_DATA PB0
#define PIN_CLK PB1
#define PIN_LATCH PB4

#define CMD_WRITE "WRITE"
#define CMD_DONE "DONE"

SoftwareSerial mySerial(PIN_RX, PIN_TX);

void latch()
{
    digitalWrite(PIN_LATCH, HIGH);
    digitalWrite(PIN_LATCH, LOW);
}

void shiftOut4(uint8_t dataPin, uint8_t clockPin, uint8_t vals[4])
{
    uint8_t vi;
    uint8_t bitOrder = MSBFIRST;

    for (vi = 0; vi < 4; vi++)
    {
        shiftOut(dataPin, clockPin, bitOrder, vals[vi]);
    }
    latch();
}

void setup()
{
    mySerial.begin(9600);

    digitalWrite(PIN_DATA, LOW);
    digitalWrite(PIN_CLK, LOW);
    digitalWrite(PIN_LATCH, LOW);

    pinMode(PIN_DATA, OUTPUT);
    pinMode(PIN_CLK, OUTPUT);
    pinMode(PIN_LATCH, OUTPUT);

    mySerial.println("EEPROM Programmer");
}

void loop()
{
    while (!mySerial.available())
        ;

    String input;
    do
    {
        input = mySerial.readStringUntil('\n');
        input.trim();
    } while (input.length() == 0);

    mySerial.println(input);
    uint8_t vals[4];

    vals[0] = strtol(input.substring(0, 2).c_str(), NULL, 16);
    vals[1] = strtol(input.substring(2, 4).c_str(), NULL, 16);
    vals[2] = strtol(input.substring(4, 6).c_str(), NULL, 16);
    vals[3] = strtol(input.substring(6, 8).c_str(), NULL, 16);

    char output2[50];
    sprintf(output2, "%02x %02x %02x %02x", vals[0], vals[1], vals[2], vals[3]);
    mySerial.println(output2);

    shiftOut4(PIN_DATA, PIN_CLK, vals);
}