#include <Arduino.h>

#define DATA_LEN 32

#define DATA_TPIN 20

#ifdef ATMEGA2560

#define CLK_INTR 21

int DATA[DATA_LEN] = {
    52,
    50,
    48,
    46,
    44,
    42,
    40,
    38,
    36, //
    53,
    51,
    49,
    47,
    45,
    43,
    41,
    39,
    37,
    35,
    33,
    31,
    29,
    27,
    25,
    23, //
    34,
    32,
    30,
    28,
    26,
    24,
    22,
};

#endif
