// Code generated from LuaParser.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // LuaParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 69, 435,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 97, 10, 3, 12,
	3, 14, 3, 100, 11, 3, 3, 3, 5, 3, 103, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 120,
	10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 151, 10, 10, 12, 10,
	14, 10, 154, 11, 10, 3, 10, 3, 10, 5, 10, 158, 10, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 178, 10, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 197, 10, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 208, 10, 17, 12, 17, 14,
	17, 211, 11, 17, 3, 18, 3, 18, 3, 18, 5, 18, 216, 10, 18, 3, 19, 3, 19,
	5, 19, 220, 10, 19, 3, 19, 5, 19, 223, 10, 19, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 7, 21, 232, 10, 21, 12, 21, 14, 21, 235, 11, 21,
	3, 21, 3, 21, 5, 21, 239, 10, 21, 3, 22, 3, 22, 3, 22, 7, 22, 244, 10,
	22, 12, 22, 14, 22, 247, 11, 22, 3, 23, 3, 23, 3, 23, 7, 23, 252, 10, 23,
	12, 23, 14, 23, 255, 11, 23, 3, 24, 3, 24, 3, 24, 7, 24, 260, 10, 24, 12,
	24, 14, 24, 263, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 278, 10, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25,
	312, 10, 25, 12, 25, 14, 25, 315, 11, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 5, 26, 323, 10, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 5, 26, 332, 10, 26, 3, 26, 3, 26, 7, 26, 336, 10, 26, 12, 26, 14,
	26, 339, 11, 26, 3, 27, 3, 27, 5, 27, 343, 10, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 5, 28, 349, 10, 28, 3, 28, 3, 28, 3, 28, 5, 28, 354, 10, 28, 3,
	29, 3, 29, 3, 29, 3, 30, 3, 30, 5, 30, 361, 10, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 370, 10, 31, 3, 31, 5, 31, 373, 10,
	31, 3, 32, 3, 32, 5, 32, 377, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 33, 7, 33, 385, 10, 33, 12, 33, 14, 33, 388, 11, 33, 3, 33, 5, 33, 391,
	10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 5, 34, 403, 10, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 417, 10, 38, 3, 39, 3, 39,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3,
	45, 3, 45, 3, 46, 3, 46, 3, 46, 2, 4, 48, 50, 47, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
	88, 90, 2, 9, 4, 2, 3, 3, 22, 22, 3, 2, 40, 41, 3, 2, 42, 45, 3, 2, 46,
	50, 5, 2, 41, 41, 48, 48, 51, 52, 3, 2, 62, 65, 3, 2, 59, 61, 2, 456, 2,
	92, 3, 2, 2, 2, 4, 98, 3, 2, 2, 2, 6, 119, 3, 2, 2, 2, 8, 121, 3, 2, 2,
	2, 10, 123, 3, 2, 2, 2, 12, 126, 3, 2, 2, 2, 14, 130, 3, 2, 2, 2, 16, 136,
	3, 2, 2, 2, 18, 141, 3, 2, 2, 2, 20, 161, 3, 2, 2, 2, 22, 169, 3, 2, 2,
	2, 24, 183, 3, 2, 2, 2, 26, 187, 3, 2, 2, 2, 28, 192, 3, 2, 2, 2, 30, 198,
	3, 2, 2, 2, 32, 202, 3, 2, 2, 2, 34, 215, 3, 2, 2, 2, 36, 217, 3, 2, 2,
	2, 38, 224, 3, 2, 2, 2, 40, 228, 3, 2, 2, 2, 42, 240, 3, 2, 2, 2, 44, 248,
	3, 2, 2, 2, 46, 256, 3, 2, 2, 2, 48, 277, 3, 2, 2, 2, 50, 322, 3, 2, 2,
	2, 52, 342, 3, 2, 2, 2, 54, 353, 3, 2, 2, 2, 56, 355, 3, 2, 2, 2, 58, 358,
	3, 2, 2, 2, 60, 372, 3, 2, 2, 2, 62, 374, 3, 2, 2, 2, 64, 380, 3, 2, 2,
	2, 66, 402, 3, 2, 2, 2, 68, 404, 3, 2, 2, 2, 70, 406, 3, 2, 2, 2, 72, 408,
	3, 2, 2, 2, 74, 416, 3, 2, 2, 2, 76, 418, 3, 2, 2, 2, 78, 420, 3, 2, 2,
	2, 80, 422, 3, 2, 2, 2, 82, 424, 3, 2, 2, 2, 84, 426, 3, 2, 2, 2, 86, 428,
	3, 2, 2, 2, 88, 430, 3, 2, 2, 2, 90, 432, 3, 2, 2, 2, 92, 93, 5, 4, 3,
	2, 93, 94, 7, 2, 2, 3, 94, 3, 3, 2, 2, 2, 95, 97, 5, 6, 4, 2, 96, 95, 3,
	2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99,
	102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 103, 5, 36, 19, 2, 102, 101,
	3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 5, 3, 2, 2, 2, 104, 120, 7, 3, 2,
	2, 105, 120, 5, 30, 16, 2, 106, 120, 5, 50, 26, 2, 107, 120, 5, 38, 20,
	2, 108, 120, 5, 8, 5, 2, 109, 120, 5, 10, 6, 2, 110, 120, 5, 12, 7, 2,
	111, 120, 5, 14, 8, 2, 112, 120, 5, 16, 9, 2, 113, 120, 5, 18, 10, 2, 114,
	120, 5, 22, 12, 2, 115, 120, 5, 20, 11, 2, 116, 120, 5, 24, 13, 2, 117,
	120, 5, 26, 14, 2, 118, 120, 5, 28, 15, 2, 119, 104, 3, 2, 2, 2, 119, 105,
	3, 2, 2, 2, 119, 106, 3, 2, 2, 2, 119, 107, 3, 2, 2, 2, 119, 108, 3, 2,
	2, 2, 119, 109, 3, 2, 2, 2, 119, 110, 3, 2, 2, 2, 119, 111, 3, 2, 2, 2,
	119, 112, 3, 2, 2, 2, 119, 113, 3, 2, 2, 2, 119, 114, 3, 2, 2, 2, 119,
	115, 3, 2, 2, 2, 119, 116, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 118,
	3, 2, 2, 2, 120, 7, 3, 2, 2, 2, 121, 122, 7, 4, 2, 2, 122, 9, 3, 2, 2,
	2, 123, 124, 7, 5, 2, 2, 124, 125, 7, 58, 2, 2, 125, 11, 3, 2, 2, 2, 126,
	127, 7, 6, 2, 2, 127, 128, 5, 4, 3, 2, 128, 129, 7, 8, 2, 2, 129, 13, 3,
	2, 2, 2, 130, 131, 7, 7, 2, 2, 131, 132, 5, 48, 25, 2, 132, 133, 7, 6,
	2, 2, 133, 134, 5, 4, 3, 2, 134, 135, 7, 8, 2, 2, 135, 15, 3, 2, 2, 2,
	136, 137, 7, 9, 2, 2, 137, 138, 5, 4, 3, 2, 138, 139, 7, 10, 2, 2, 139,
	140, 5, 48, 25, 2, 140, 17, 3, 2, 2, 2, 141, 142, 7, 14, 2, 2, 142, 143,
	5, 48, 25, 2, 143, 144, 7, 15, 2, 2, 144, 152, 5, 4, 3, 2, 145, 146, 7,
	16, 2, 2, 146, 147, 5, 48, 25, 2, 147, 148, 7, 15, 2, 2, 148, 149, 5, 4,
	3, 2, 149, 151, 3, 2, 2, 2, 150, 145, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2,
	152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 157, 3, 2, 2, 2, 154,
	152, 3, 2, 2, 2, 155, 156, 7, 17, 2, 2, 156, 158, 5, 4, 3, 2, 157, 155,
	3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 7, 8,
	2, 2, 160, 19, 3, 2, 2, 2, 161, 162, 7, 11, 2, 2, 162, 163, 5, 44, 23,
	2, 163, 164, 7, 23, 2, 2, 164, 165, 5, 46, 24, 2, 165, 166, 7, 6, 2, 2,
	166, 167, 5, 4, 3, 2, 167, 168, 7, 8, 2, 2, 168, 21, 3, 2, 2, 2, 169, 170,
	7, 11, 2, 2, 170, 171, 7, 58, 2, 2, 171, 172, 7, 38, 2, 2, 172, 173, 5,
	48, 25, 2, 173, 174, 7, 22, 2, 2, 174, 177, 5, 48, 25, 2, 175, 176, 7,
	22, 2, 2, 176, 178, 5, 48, 25, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2,
	2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 7, 6, 2, 2, 180, 181, 5, 4, 3, 2,
	181, 182, 7, 8, 2, 2, 182, 23, 3, 2, 2, 2, 183, 184, 7, 12, 2, 2, 184,
	185, 5, 40, 21, 2, 185, 186, 5, 58, 30, 2, 186, 25, 3, 2, 2, 2, 187, 188,
	7, 13, 2, 2, 188, 189, 7, 12, 2, 2, 189, 190, 7, 58, 2, 2, 190, 191, 5,
	58, 30, 2, 191, 27, 3, 2, 2, 2, 192, 193, 7, 13, 2, 2, 193, 196, 5, 32,
	17, 2, 194, 195, 7, 38, 2, 2, 195, 197, 5, 46, 24, 2, 196, 194, 3, 2, 2,
	2, 196, 197, 3, 2, 2, 2, 197, 29, 3, 2, 2, 2, 198, 199, 5, 42, 22, 2, 199,
	200, 7, 38, 2, 2, 200, 201, 5, 46, 24, 2, 201, 31, 3, 2, 2, 2, 202, 203,
	7, 58, 2, 2, 203, 209, 5, 34, 18, 2, 204, 205, 7, 22, 2, 2, 205, 206, 7,
	58, 2, 2, 206, 208, 5, 34, 18, 2, 207, 204, 3, 2, 2, 2, 208, 211, 3, 2,
	2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 33, 3, 2, 2, 2,
	211, 209, 3, 2, 2, 2, 212, 213, 7, 32, 2, 2, 213, 214, 7, 58, 2, 2, 214,
	216, 7, 33, 2, 2, 215, 212, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 35,
	3, 2, 2, 2, 217, 219, 7, 18, 2, 2, 218, 220, 5, 46, 24, 2, 219, 218, 3,
	2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 222, 3, 2, 2, 2, 221, 223, 7, 3, 2,
	2, 222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 37, 3, 2, 2, 2, 224,
	225, 7, 20, 2, 2, 225, 226, 7, 58, 2, 2, 226, 227, 7, 20, 2, 2, 227, 39,
	3, 2, 2, 2, 228, 233, 7, 58, 2, 2, 229, 230, 7, 21, 2, 2, 230, 232, 7,
	58, 2, 2, 231, 229, 3, 2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 231, 3, 2, 2,
	2, 233, 234, 3, 2, 2, 2, 234, 238, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 236,
	237, 7, 19, 2, 2, 237, 239, 7, 58, 2, 2, 238, 236, 3, 2, 2, 2, 238, 239,
	3, 2, 2, 2, 239, 41, 3, 2, 2, 2, 240, 245, 5, 50, 26, 2, 241, 242, 7, 22,
	2, 2, 242, 244, 5, 50, 26, 2, 243, 241, 3, 2, 2, 2, 244, 247, 3, 2, 2,
	2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 43, 3, 2, 2, 2, 247,
	245, 3, 2, 2, 2, 248, 253, 7, 58, 2, 2, 249, 250, 7, 22, 2, 2, 250, 252,
	7, 58, 2, 2, 251, 249, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2,
	2, 2, 253, 254, 3, 2, 2, 2, 254, 45, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2,
	256, 261, 5, 48, 25, 2, 257, 258, 7, 22, 2, 2, 258, 260, 5, 48, 25, 2,
	259, 257, 3, 2, 2, 2, 260, 263, 3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261,
	262, 3, 2, 2, 2, 262, 47, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 264, 265, 8,
	25, 1, 2, 265, 278, 7, 54, 2, 2, 266, 278, 7, 55, 2, 2, 267, 278, 7, 56,
	2, 2, 268, 278, 5, 88, 45, 2, 269, 278, 5, 90, 46, 2, 270, 278, 7, 57,
	2, 2, 271, 278, 5, 56, 29, 2, 272, 278, 5, 50, 26, 2, 273, 278, 5, 62,
	32, 2, 274, 275, 5, 84, 43, 2, 275, 276, 5, 48, 25, 10, 276, 278, 3, 2,
	2, 2, 277, 264, 3, 2, 2, 2, 277, 266, 3, 2, 2, 2, 277, 267, 3, 2, 2, 2,
	277, 268, 3, 2, 2, 2, 277, 269, 3, 2, 2, 2, 277, 270, 3, 2, 2, 2, 277,
	271, 3, 2, 2, 2, 277, 272, 3, 2, 2, 2, 277, 273, 3, 2, 2, 2, 277, 274,
	3, 2, 2, 2, 278, 313, 3, 2, 2, 2, 279, 280, 12, 11, 2, 2, 280, 281, 5,
	86, 44, 2, 281, 282, 5, 48, 25, 11, 282, 312, 3, 2, 2, 2, 283, 284, 12,
	9, 2, 2, 284, 285, 5, 80, 41, 2, 285, 286, 5, 48, 25, 10, 286, 312, 3,
	2, 2, 2, 287, 288, 12, 8, 2, 2, 288, 289, 5, 78, 40, 2, 289, 290, 5, 48,
	25, 9, 290, 312, 3, 2, 2, 2, 291, 292, 12, 7, 2, 2, 292, 293, 5, 76, 39,
	2, 293, 294, 5, 48, 25, 7, 294, 312, 3, 2, 2, 2, 295, 296, 12, 6, 2, 2,
	296, 297, 5, 74, 38, 2, 297, 298, 5, 48, 25, 7, 298, 312, 3, 2, 2, 2, 299,
	300, 12, 5, 2, 2, 300, 301, 5, 72, 37, 2, 301, 302, 5, 48, 25, 6, 302,
	312, 3, 2, 2, 2, 303, 304, 12, 4, 2, 2, 304, 305, 5, 70, 36, 2, 305, 306,
	5, 48, 25, 5, 306, 312, 3, 2, 2, 2, 307, 308, 12, 3, 2, 2, 308, 309, 5,
	82, 42, 2, 309, 310, 5, 48, 25, 4, 310, 312, 3, 2, 2, 2, 311, 279, 3, 2,
	2, 2, 311, 283, 3, 2, 2, 2, 311, 287, 3, 2, 2, 2, 311, 291, 3, 2, 2, 2,
	311, 295, 3, 2, 2, 2, 311, 299, 3, 2, 2, 2, 311, 303, 3, 2, 2, 2, 311,
	307, 3, 2, 2, 2, 312, 315, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314,
	3, 2, 2, 2, 314, 49, 3, 2, 2, 2, 315, 313, 3, 2, 2, 2, 316, 317, 8, 26,
	1, 2, 317, 323, 7, 58, 2, 2, 318, 319, 7, 24, 2, 2, 319, 320, 5, 48, 25,
	2, 320, 321, 7, 25, 2, 2, 321, 323, 3, 2, 2, 2, 322, 316, 3, 2, 2, 2, 322,
	318, 3, 2, 2, 2, 323, 337, 3, 2, 2, 2, 324, 331, 12, 5, 2, 2, 325, 326,
	7, 26, 2, 2, 326, 327, 5, 48, 25, 2, 327, 328, 7, 27, 2, 2, 328, 332, 3,
	2, 2, 2, 329, 330, 7, 21, 2, 2, 330, 332, 7, 58, 2, 2, 331, 325, 3, 2,
	2, 2, 331, 329, 3, 2, 2, 2, 332, 336, 3, 2, 2, 2, 333, 334, 12, 3, 2, 2,
	334, 336, 5, 52, 27, 2, 335, 324, 3, 2, 2, 2, 335, 333, 3, 2, 2, 2, 336,
	339, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 51, 3,
	2, 2, 2, 339, 337, 3, 2, 2, 2, 340, 341, 7, 19, 2, 2, 341, 343, 7, 58,
	2, 2, 342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2,
	344, 345, 5, 54, 28, 2, 345, 53, 3, 2, 2, 2, 346, 348, 7, 24, 2, 2, 347,
	349, 5, 46, 24, 2, 348, 347, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 350,
	3, 2, 2, 2, 350, 354, 7, 25, 2, 2, 351, 354, 5, 62, 32, 2, 352, 354, 5,
	90, 46, 2, 353, 346, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 353, 352, 3, 2,
	2, 2, 354, 55, 3, 2, 2, 2, 355, 356, 7, 12, 2, 2, 356, 357, 5, 58, 30,
	2, 357, 57, 3, 2, 2, 2, 358, 360, 7, 24, 2, 2, 359, 361, 5, 60, 31, 2,
	360, 359, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362,
	363, 7, 25, 2, 2, 363, 364, 5, 4, 3, 2, 364, 365, 7, 8, 2, 2, 365, 59,
	3, 2, 2, 2, 366, 369, 5, 44, 23, 2, 367, 368, 7, 22, 2, 2, 368, 370, 7,
	57, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 373, 3, 2, 2,
	2, 371, 373, 7, 57, 2, 2, 372, 366, 3, 2, 2, 2, 372, 371, 3, 2, 2, 2, 373,
	61, 3, 2, 2, 2, 374, 376, 7, 28, 2, 2, 375, 377, 5, 64, 33, 2, 376, 375,
	3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 379, 7, 29,
	2, 2, 379, 63, 3, 2, 2, 2, 380, 386, 5, 66, 34, 2, 381, 382, 5, 68, 35,
	2, 382, 383, 5, 66, 34, 2, 383, 385, 3, 2, 2, 2, 384, 381, 3, 2, 2, 2,
	385, 388, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387,
	390, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 389, 391, 5, 68, 35, 2, 390, 389,
	3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 65, 3, 2, 2, 2, 392, 393, 7, 26,
	2, 2, 393, 394, 5, 48, 25, 2, 394, 395, 7, 27, 2, 2, 395, 396, 7, 38, 2,
	2, 396, 397, 5, 48, 25, 2, 397, 403, 3, 2, 2, 2, 398, 399, 7, 58, 2, 2,
	399, 400, 7, 38, 2, 2, 400, 403, 5, 48, 25, 2, 401, 403, 5, 48, 25, 2,
	402, 392, 3, 2, 2, 2, 402, 398, 3, 2, 2, 2, 402, 401, 3, 2, 2, 2, 403,
	67, 3, 2, 2, 2, 404, 405, 9, 2, 2, 2, 405, 69, 3, 2, 2, 2, 406, 407, 7,
	30, 2, 2, 407, 71, 3, 2, 2, 2, 408, 409, 7, 31, 2, 2, 409, 73, 3, 2, 2,
	2, 410, 417, 7, 32, 2, 2, 411, 417, 7, 33, 2, 2, 412, 417, 7, 34, 2, 2,
	413, 417, 7, 35, 2, 2, 414, 417, 7, 36, 2, 2, 415, 417, 7, 37, 2, 2, 416,
	410, 3, 2, 2, 2, 416, 411, 3, 2, 2, 2, 416, 412, 3, 2, 2, 2, 416, 413,
	3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 416, 415, 3, 2, 2, 2, 417, 75, 3, 2,
	2, 2, 418, 419, 7, 39, 2, 2, 419, 77, 3, 2, 2, 2, 420, 421, 9, 3, 2, 2,
	421, 79, 3, 2, 2, 2, 422, 423, 9, 4, 2, 2, 423, 81, 3, 2, 2, 2, 424, 425,
	9, 5, 2, 2, 425, 83, 3, 2, 2, 2, 426, 427, 9, 6, 2, 2, 427, 85, 3, 2, 2,
	2, 428, 429, 7, 53, 2, 2, 429, 87, 3, 2, 2, 2, 430, 431, 9, 7, 2, 2, 431,
	89, 3, 2, 2, 2, 432, 433, 9, 8, 2, 2, 433, 91, 3, 2, 2, 2, 36, 98, 102,
	119, 152, 157, 177, 196, 209, 215, 219, 222, 233, 238, 245, 253, 261, 277,
	311, 313, 322, 331, 335, 337, 342, 348, 353, 360, 369, 372, 376, 386, 390,
	402, 416,
}
var literalNames = []string{
	"", "';'", "'break'", "'goto'", "'do'", "'while'", "'end'", "'repeat'",
	"'until'", "'for'", "'function'", "'local'", "'if'", "'then'", "'elseif'",
	"'else'", "'return'", "':'", "'::'", "'.'", "','", "'in'", "'('", "')'",
	"'['", "']'", "'{'", "'}'", "'or'", "'and'", "'<'", "'>'", "'<='", "'>='",
	"'~='", "'=='", "'='", "'..'", "'+'", "'-'", "'*'", "'/'", "'%'", "'//'",
	"'&'", "'|'", "'~'", "'<<'", "'>>'", "'not'", "'#'", "'^'", "'nil'", "'false'",
	"'true'", "'...'",
}
var symbolicNames = []string{
	"", "SEMICOLON", "BREAK", "GOTO", "DO", "WHILE", "END", "REPEAT", "UNTIL",
	"FOR", "FUNCTION", "LOCAL", "IF", "THEN", "ELSEIF", "ELSE", "RETURN", "COLON",
	"DCOLON", "DOT", "COMMA", "IN", "LPAREN", "RPAREN", "LBRACK", "RBRACK",
	"LBRACE", "RBRACE", "OR", "AND", "LT", "GT", "LTE", "GTE", "NEQ", "EQ",
	"EQUALS", "STRCAT", "PLUS", "MINUS", "MUL", "DIV", "MOD", "DIVFLOOR", "BITAND",
	"BITOR", "BITNOT", "BITSHL", "BITSHR", "NOT", "LEN", "POWER", "NIL", "FALSE",
	"TRUE", "DOTS", "NAME", "NORMALSTRING", "CHARSTRING", "LONGSTRING", "INT",
	"HEX", "FLOAT", "HEX_FLOAT", "COMMENT", "LINE_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"chunk", "block", "stat", "breakstat", "gotostat", "dostat", "whilestat",
	"repeatstat", "ifstat", "genericforstat", "numericforstat", "functionstat",
	"localfunctionstat", "localvariablestat", "variablestat", "attnamelist",
	"attrib", "retstat", "label", "funcname", "variablelist", "namelist", "explist",
	"exp", "variable", "nameAndArgs", "args", "functiondef", "funcbody", "parlist",
	"tableconstructor", "fieldlist", "field", "fieldsep", "operatorOr", "operatorAnd",
	"operatorComparison", "operatorStrcat", "operatorAddSub", "operatorMulDivMod",
	"operatorBitwise", "operatorUnary", "operatorPower", "number", "lstring",
}

type LuaParser struct {
	*antlr.BaseParser
}

// NewLuaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LuaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLuaParser(input antlr.TokenStream) *LuaParser {
	this := new(LuaParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LuaParser.g4"

	return this
}

// LuaParser tokens.
const (
	LuaParserEOF          = antlr.TokenEOF
	LuaParserSEMICOLON    = 1
	LuaParserBREAK        = 2
	LuaParserGOTO         = 3
	LuaParserDO           = 4
	LuaParserWHILE        = 5
	LuaParserEND          = 6
	LuaParserREPEAT       = 7
	LuaParserUNTIL        = 8
	LuaParserFOR          = 9
	LuaParserFUNCTION     = 10
	LuaParserLOCAL        = 11
	LuaParserIF           = 12
	LuaParserTHEN         = 13
	LuaParserELSEIF       = 14
	LuaParserELSE         = 15
	LuaParserRETURN       = 16
	LuaParserCOLON        = 17
	LuaParserDCOLON       = 18
	LuaParserDOT          = 19
	LuaParserCOMMA        = 20
	LuaParserIN           = 21
	LuaParserLPAREN       = 22
	LuaParserRPAREN       = 23
	LuaParserLBRACK       = 24
	LuaParserRBRACK       = 25
	LuaParserLBRACE       = 26
	LuaParserRBRACE       = 27
	LuaParserOR           = 28
	LuaParserAND          = 29
	LuaParserLT           = 30
	LuaParserGT           = 31
	LuaParserLTE          = 32
	LuaParserGTE          = 33
	LuaParserNEQ          = 34
	LuaParserEQ           = 35
	LuaParserEQUALS       = 36
	LuaParserSTRCAT       = 37
	LuaParserPLUS         = 38
	LuaParserMINUS        = 39
	LuaParserMUL          = 40
	LuaParserDIV          = 41
	LuaParserMOD          = 42
	LuaParserDIVFLOOR     = 43
	LuaParserBITAND       = 44
	LuaParserBITOR        = 45
	LuaParserBITNOT       = 46
	LuaParserBITSHL       = 47
	LuaParserBITSHR       = 48
	LuaParserNOT          = 49
	LuaParserLEN          = 50
	LuaParserPOWER        = 51
	LuaParserNIL          = 52
	LuaParserFALSE        = 53
	LuaParserTRUE         = 54
	LuaParserDOTS         = 55
	LuaParserNAME         = 56
	LuaParserNORMALSTRING = 57
	LuaParserCHARSTRING   = 58
	LuaParserLONGSTRING   = 59
	LuaParserINT          = 60
	LuaParserHEX          = 61
	LuaParserFLOAT        = 62
	LuaParserHEX_FLOAT    = 63
	LuaParserCOMMENT      = 64
	LuaParserLINE_COMMENT = 65
	LuaParserWS           = 66
	LuaParserSHEBANG      = 67
)

// LuaParser rules.
const (
	LuaParserRULE_chunk              = 0
	LuaParserRULE_block              = 1
	LuaParserRULE_stat               = 2
	LuaParserRULE_breakstat          = 3
	LuaParserRULE_gotostat           = 4
	LuaParserRULE_dostat             = 5
	LuaParserRULE_whilestat          = 6
	LuaParserRULE_repeatstat         = 7
	LuaParserRULE_ifstat             = 8
	LuaParserRULE_genericforstat     = 9
	LuaParserRULE_numericforstat     = 10
	LuaParserRULE_functionstat       = 11
	LuaParserRULE_localfunctionstat  = 12
	LuaParserRULE_localvariablestat  = 13
	LuaParserRULE_variablestat       = 14
	LuaParserRULE_attnamelist        = 15
	LuaParserRULE_attrib             = 16
	LuaParserRULE_retstat            = 17
	LuaParserRULE_label              = 18
	LuaParserRULE_funcname           = 19
	LuaParserRULE_variablelist       = 20
	LuaParserRULE_namelist           = 21
	LuaParserRULE_explist            = 22
	LuaParserRULE_exp                = 23
	LuaParserRULE_variable           = 24
	LuaParserRULE_nameAndArgs        = 25
	LuaParserRULE_args               = 26
	LuaParserRULE_functiondef        = 27
	LuaParserRULE_funcbody           = 28
	LuaParserRULE_parlist            = 29
	LuaParserRULE_tableconstructor   = 30
	LuaParserRULE_fieldlist          = 31
	LuaParserRULE_field              = 32
	LuaParserRULE_fieldsep           = 33
	LuaParserRULE_operatorOr         = 34
	LuaParserRULE_operatorAnd        = 35
	LuaParserRULE_operatorComparison = 36
	LuaParserRULE_operatorStrcat     = 37
	LuaParserRULE_operatorAddSub     = 38
	LuaParserRULE_operatorMulDivMod  = 39
	LuaParserRULE_operatorBitwise    = 40
	LuaParserRULE_operatorUnary      = 41
	LuaParserRULE_operatorPower      = 42
	LuaParserRULE_number             = 43
	LuaParserRULE_lstring            = 44
)

// IChunkContext is an interface to support dynamic dispatch.
type IChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChunkContext differentiates from other interfaces.
	IsChunkContext()
}

type ChunkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChunkContext() *ChunkContext {
	var p = new(ChunkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_chunk
	return p
}

func (*ChunkContext) IsChunkContext() {}

func NewChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChunkContext {
	var p = new(ChunkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_chunk

	return p
}

func (s *ChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *ChunkContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ChunkContext) EOF() antlr.TerminalNode {
	return s.GetToken(LuaParserEOF, 0)
}

func (s *ChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChunkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterChunk(s)
	}
}

func (s *ChunkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitChunk(s)
	}
}

func (p *LuaParser) Chunk() (localctx IChunkContext) {
	localctx = NewChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LuaParserRULE_chunk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Block()
	}
	{
		p.SetState(91)
		p.Match(LuaParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) Retstat() IRetstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetstatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *LuaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LuaParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserSEMICOLON)|(1<<LuaParserBREAK)|(1<<LuaParserGOTO)|(1<<LuaParserDO)|(1<<LuaParserWHILE)|(1<<LuaParserREPEAT)|(1<<LuaParserFOR)|(1<<LuaParserFUNCTION)|(1<<LuaParserLOCAL)|(1<<LuaParserIF)|(1<<LuaParserDCOLON)|(1<<LuaParserLPAREN))) != 0) || _la == LuaParserNAME {
		{
			p.SetState(93)
			p.Stat()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserRETURN {
		{
			p.SetState(99)
			p.Retstat()
		}

	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LuaParserSEMICOLON, 0)
}

func (s *StatContext) Variablestat() IVariablestatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablestatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablestatContext)
}

func (s *StatContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *StatContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *StatContext) Breakstat() IBreakstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakstatContext)
}

func (s *StatContext) Gotostat() IGotostatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotostatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotostatContext)
}

func (s *StatContext) Dostat() IDostatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDostatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDostatContext)
}

func (s *StatContext) Whilestat() IWhilestatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhilestatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhilestatContext)
}

func (s *StatContext) Repeatstat() IRepeatstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepeatstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepeatstatContext)
}

func (s *StatContext) Ifstat() IIfstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatContext)
}

func (s *StatContext) Numericforstat() INumericforstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericforstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericforstatContext)
}

func (s *StatContext) Genericforstat() IGenericforstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGenericforstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGenericforstatContext)
}

func (s *StatContext) Functionstat() IFunctionstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionstatContext)
}

func (s *StatContext) Localfunctionstat() ILocalfunctionstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalfunctionstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalfunctionstatContext)
}

func (s *StatContext) Localvariablestat() ILocalvariablestatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalvariablestatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalvariablestatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *LuaParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LuaParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Match(LuaParserSEMICOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Variablestat()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.variable(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Label()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
			p.Breakstat()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(107)
			p.Gotostat()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(108)
			p.Dostat()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(109)
			p.Whilestat()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(110)
			p.Repeatstat()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(111)
			p.Ifstat()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(112)
			p.Numericforstat()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(113)
			p.Genericforstat()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(114)
			p.Functionstat()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(115)
			p.Localfunctionstat()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(116)
			p.Localvariablestat()
		}

	}

	return localctx
}

// IBreakstatContext is an interface to support dynamic dispatch.
type IBreakstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakstatContext differentiates from other interfaces.
	IsBreakstatContext()
}

type BreakstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakstatContext() *BreakstatContext {
	var p = new(BreakstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_breakstat
	return p
}

func (*BreakstatContext) IsBreakstatContext() {}

func NewBreakstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakstatContext {
	var p = new(BreakstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_breakstat

	return p
}

func (s *BreakstatContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakstatContext) BREAK() antlr.TerminalNode {
	return s.GetToken(LuaParserBREAK, 0)
}

func (s *BreakstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterBreakstat(s)
	}
}

func (s *BreakstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitBreakstat(s)
	}
}

func (p *LuaParser) Breakstat() (localctx IBreakstatContext) {
	localctx = NewBreakstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LuaParserRULE_breakstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(LuaParserBREAK)
	}

	return localctx
}

// IGotostatContext is an interface to support dynamic dispatch.
type IGotostatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotostatContext differentiates from other interfaces.
	IsGotostatContext()
}

type GotostatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotostatContext() *GotostatContext {
	var p = new(GotostatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_gotostat
	return p
}

func (*GotostatContext) IsGotostatContext() {}

func NewGotostatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotostatContext {
	var p = new(GotostatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_gotostat

	return p
}

func (s *GotostatContext) GetParser() antlr.Parser { return s.parser }

func (s *GotostatContext) GOTO() antlr.TerminalNode {
	return s.GetToken(LuaParserGOTO, 0)
}

func (s *GotostatContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *GotostatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotostatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotostatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterGotostat(s)
	}
}

func (s *GotostatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitGotostat(s)
	}
}

func (p *LuaParser) Gotostat() (localctx IGotostatContext) {
	localctx = NewGotostatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LuaParserRULE_gotostat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(LuaParserGOTO)
	}
	{
		p.SetState(122)
		p.Match(LuaParserNAME)
	}

	return localctx
}

// IDostatContext is an interface to support dynamic dispatch.
type IDostatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDostatContext differentiates from other interfaces.
	IsDostatContext()
}

type DostatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDostatContext() *DostatContext {
	var p = new(DostatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_dostat
	return p
}

func (*DostatContext) IsDostatContext() {}

func NewDostatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DostatContext {
	var p = new(DostatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_dostat

	return p
}

func (s *DostatContext) GetParser() antlr.Parser { return s.parser }

func (s *DostatContext) DO() antlr.TerminalNode {
	return s.GetToken(LuaParserDO, 0)
}

func (s *DostatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DostatContext) END() antlr.TerminalNode {
	return s.GetToken(LuaParserEND, 0)
}

func (s *DostatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DostatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DostatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterDostat(s)
	}
}

func (s *DostatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitDostat(s)
	}
}

func (p *LuaParser) Dostat() (localctx IDostatContext) {
	localctx = NewDostatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LuaParserRULE_dostat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(LuaParserDO)
	}
	{
		p.SetState(125)
		p.Block()
	}
	{
		p.SetState(126)
		p.Match(LuaParserEND)
	}

	return localctx
}

// IWhilestatContext is an interface to support dynamic dispatch.
type IWhilestatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhilestatContext differentiates from other interfaces.
	IsWhilestatContext()
}

type WhilestatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestatContext() *WhilestatContext {
	var p = new(WhilestatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_whilestat
	return p
}

func (*WhilestatContext) IsWhilestatContext() {}

func NewWhilestatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestatContext {
	var p = new(WhilestatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_whilestat

	return p
}

func (s *WhilestatContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestatContext) WHILE() antlr.TerminalNode {
	return s.GetToken(LuaParserWHILE, 0)
}

func (s *WhilestatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *WhilestatContext) DO() antlr.TerminalNode {
	return s.GetToken(LuaParserDO, 0)
}

func (s *WhilestatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestatContext) END() antlr.TerminalNode {
	return s.GetToken(LuaParserEND, 0)
}

func (s *WhilestatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterWhilestat(s)
	}
}

func (s *WhilestatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitWhilestat(s)
	}
}

func (p *LuaParser) Whilestat() (localctx IWhilestatContext) {
	localctx = NewWhilestatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LuaParserRULE_whilestat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(LuaParserWHILE)
	}
	{
		p.SetState(129)
		p.exp(0)
	}
	{
		p.SetState(130)
		p.Match(LuaParserDO)
	}
	{
		p.SetState(131)
		p.Block()
	}
	{
		p.SetState(132)
		p.Match(LuaParserEND)
	}

	return localctx
}

// IRepeatstatContext is an interface to support dynamic dispatch.
type IRepeatstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepeatstatContext differentiates from other interfaces.
	IsRepeatstatContext()
}

type RepeatstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatstatContext() *RepeatstatContext {
	var p = new(RepeatstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_repeatstat
	return p
}

func (*RepeatstatContext) IsRepeatstatContext() {}

func NewRepeatstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatstatContext {
	var p = new(RepeatstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_repeatstat

	return p
}

func (s *RepeatstatContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatstatContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(LuaParserREPEAT, 0)
}

func (s *RepeatstatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RepeatstatContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(LuaParserUNTIL, 0)
}

func (s *RepeatstatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *RepeatstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterRepeatstat(s)
	}
}

func (s *RepeatstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitRepeatstat(s)
	}
}

func (p *LuaParser) Repeatstat() (localctx IRepeatstatContext) {
	localctx = NewRepeatstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LuaParserRULE_repeatstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(LuaParserREPEAT)
	}
	{
		p.SetState(135)
		p.Block()
	}
	{
		p.SetState(136)
		p.Match(LuaParserUNTIL)
	}
	{
		p.SetState(137)
		p.exp(0)
	}

	return localctx
}

// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) IF() antlr.TerminalNode {
	return s.GetToken(LuaParserIF, 0)
}

func (s *IfstatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *IfstatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IfstatContext) AllTHEN() []antlr.TerminalNode {
	return s.GetTokens(LuaParserTHEN)
}

func (s *IfstatContext) THEN(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserTHEN, i)
}

func (s *IfstatContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfstatContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstatContext) END() antlr.TerminalNode {
	return s.GetToken(LuaParserEND, 0)
}

func (s *IfstatContext) AllELSEIF() []antlr.TerminalNode {
	return s.GetTokens(LuaParserELSEIF)
}

func (s *IfstatContext) ELSEIF(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserELSEIF, i)
}

func (s *IfstatContext) ELSE() antlr.TerminalNode {
	return s.GetToken(LuaParserELSE, 0)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterIfstat(s)
	}
}

func (s *IfstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitIfstat(s)
	}
}

func (p *LuaParser) Ifstat() (localctx IIfstatContext) {
	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LuaParserRULE_ifstat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(LuaParserIF)
	}
	{
		p.SetState(140)
		p.exp(0)
	}
	{
		p.SetState(141)
		p.Match(LuaParserTHEN)
	}
	{
		p.SetState(142)
		p.Block()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserELSEIF {
		{
			p.SetState(143)
			p.Match(LuaParserELSEIF)
		}
		{
			p.SetState(144)
			p.exp(0)
		}
		{
			p.SetState(145)
			p.Match(LuaParserTHEN)
		}
		{
			p.SetState(146)
			p.Block()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserELSE {
		{
			p.SetState(153)
			p.Match(LuaParserELSE)
		}
		{
			p.SetState(154)
			p.Block()
		}

	}
	{
		p.SetState(157)
		p.Match(LuaParserEND)
	}

	return localctx
}

// IGenericforstatContext is an interface to support dynamic dispatch.
type IGenericforstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericforstatContext differentiates from other interfaces.
	IsGenericforstatContext()
}

type GenericforstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericforstatContext() *GenericforstatContext {
	var p = new(GenericforstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_genericforstat
	return p
}

func (*GenericforstatContext) IsGenericforstatContext() {}

func NewGenericforstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericforstatContext {
	var p = new(GenericforstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_genericforstat

	return p
}

func (s *GenericforstatContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericforstatContext) FOR() antlr.TerminalNode {
	return s.GetToken(LuaParserFOR, 0)
}

func (s *GenericforstatContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *GenericforstatContext) IN() antlr.TerminalNode {
	return s.GetToken(LuaParserIN, 0)
}

func (s *GenericforstatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *GenericforstatContext) DO() antlr.TerminalNode {
	return s.GetToken(LuaParserDO, 0)
}

func (s *GenericforstatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GenericforstatContext) END() antlr.TerminalNode {
	return s.GetToken(LuaParserEND, 0)
}

func (s *GenericforstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericforstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericforstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterGenericforstat(s)
	}
}

func (s *GenericforstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitGenericforstat(s)
	}
}

func (p *LuaParser) Genericforstat() (localctx IGenericforstatContext) {
	localctx = NewGenericforstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LuaParserRULE_genericforstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(LuaParserFOR)
	}
	{
		p.SetState(160)
		p.Namelist()
	}
	{
		p.SetState(161)
		p.Match(LuaParserIN)
	}
	{
		p.SetState(162)
		p.Explist()
	}
	{
		p.SetState(163)
		p.Match(LuaParserDO)
	}
	{
		p.SetState(164)
		p.Block()
	}
	{
		p.SetState(165)
		p.Match(LuaParserEND)
	}

	return localctx
}

// INumericforstatContext is an interface to support dynamic dispatch.
type INumericforstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericforstatContext differentiates from other interfaces.
	IsNumericforstatContext()
}

type NumericforstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericforstatContext() *NumericforstatContext {
	var p = new(NumericforstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_numericforstat
	return p
}

func (*NumericforstatContext) IsNumericforstatContext() {}

func NewNumericforstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericforstatContext {
	var p = new(NumericforstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_numericforstat

	return p
}

func (s *NumericforstatContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericforstatContext) FOR() antlr.TerminalNode {
	return s.GetToken(LuaParserFOR, 0)
}

func (s *NumericforstatContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *NumericforstatContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(LuaParserEQUALS, 0)
}

func (s *NumericforstatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *NumericforstatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *NumericforstatContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LuaParserCOMMA)
}

func (s *NumericforstatContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserCOMMA, i)
}

func (s *NumericforstatContext) DO() antlr.TerminalNode {
	return s.GetToken(LuaParserDO, 0)
}

func (s *NumericforstatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *NumericforstatContext) END() antlr.TerminalNode {
	return s.GetToken(LuaParserEND, 0)
}

func (s *NumericforstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericforstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericforstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterNumericforstat(s)
	}
}

func (s *NumericforstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitNumericforstat(s)
	}
}

func (p *LuaParser) Numericforstat() (localctx INumericforstatContext) {
	localctx = NewNumericforstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LuaParserRULE_numericforstat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(LuaParserFOR)
	}
	{
		p.SetState(168)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(169)
		p.Match(LuaParserEQUALS)
	}
	{
		p.SetState(170)
		p.exp(0)
	}
	{
		p.SetState(171)
		p.Match(LuaParserCOMMA)
	}
	{
		p.SetState(172)
		p.exp(0)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserCOMMA {
		{
			p.SetState(173)
			p.Match(LuaParserCOMMA)
		}
		{
			p.SetState(174)
			p.exp(0)
		}

	}
	{
		p.SetState(177)
		p.Match(LuaParserDO)
	}
	{
		p.SetState(178)
		p.Block()
	}
	{
		p.SetState(179)
		p.Match(LuaParserEND)
	}

	return localctx
}

// IFunctionstatContext is an interface to support dynamic dispatch.
type IFunctionstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionstatContext differentiates from other interfaces.
	IsFunctionstatContext()
}

type FunctionstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionstatContext() *FunctionstatContext {
	var p = new(FunctionstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_functionstat
	return p
}

func (*FunctionstatContext) IsFunctionstatContext() {}

func NewFunctionstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionstatContext {
	var p = new(FunctionstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_functionstat

	return p
}

func (s *FunctionstatContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionstatContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(LuaParserFUNCTION, 0)
}

func (s *FunctionstatContext) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *FunctionstatContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FunctionstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterFunctionstat(s)
	}
}

func (s *FunctionstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitFunctionstat(s)
	}
}

func (p *LuaParser) Functionstat() (localctx IFunctionstatContext) {
	localctx = NewFunctionstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LuaParserRULE_functionstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(LuaParserFUNCTION)
	}
	{
		p.SetState(182)
		p.Funcname()
	}
	{
		p.SetState(183)
		p.Funcbody()
	}

	return localctx
}

// ILocalfunctionstatContext is an interface to support dynamic dispatch.
type ILocalfunctionstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalfunctionstatContext differentiates from other interfaces.
	IsLocalfunctionstatContext()
}

type LocalfunctionstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalfunctionstatContext() *LocalfunctionstatContext {
	var p = new(LocalfunctionstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_localfunctionstat
	return p
}

func (*LocalfunctionstatContext) IsLocalfunctionstatContext() {}

func NewLocalfunctionstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalfunctionstatContext {
	var p = new(LocalfunctionstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_localfunctionstat

	return p
}

func (s *LocalfunctionstatContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalfunctionstatContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(LuaParserLOCAL, 0)
}

func (s *LocalfunctionstatContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(LuaParserFUNCTION, 0)
}

func (s *LocalfunctionstatContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *LocalfunctionstatContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *LocalfunctionstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalfunctionstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalfunctionstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterLocalfunctionstat(s)
	}
}

func (s *LocalfunctionstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitLocalfunctionstat(s)
	}
}

func (p *LuaParser) Localfunctionstat() (localctx ILocalfunctionstatContext) {
	localctx = NewLocalfunctionstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LuaParserRULE_localfunctionstat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(LuaParserLOCAL)
	}
	{
		p.SetState(186)
		p.Match(LuaParserFUNCTION)
	}
	{
		p.SetState(187)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(188)
		p.Funcbody()
	}

	return localctx
}

// ILocalvariablestatContext is an interface to support dynamic dispatch.
type ILocalvariablestatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalvariablestatContext differentiates from other interfaces.
	IsLocalvariablestatContext()
}

type LocalvariablestatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalvariablestatContext() *LocalvariablestatContext {
	var p = new(LocalvariablestatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_localvariablestat
	return p
}

func (*LocalvariablestatContext) IsLocalvariablestatContext() {}

func NewLocalvariablestatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalvariablestatContext {
	var p = new(LocalvariablestatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_localvariablestat

	return p
}

func (s *LocalvariablestatContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalvariablestatContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(LuaParserLOCAL, 0)
}

func (s *LocalvariablestatContext) Attnamelist() IAttnamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttnamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttnamelistContext)
}

func (s *LocalvariablestatContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(LuaParserEQUALS, 0)
}

func (s *LocalvariablestatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *LocalvariablestatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalvariablestatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalvariablestatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterLocalvariablestat(s)
	}
}

func (s *LocalvariablestatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitLocalvariablestat(s)
	}
}

func (p *LuaParser) Localvariablestat() (localctx ILocalvariablestatContext) {
	localctx = NewLocalvariablestatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LuaParserRULE_localvariablestat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(LuaParserLOCAL)
	}
	{
		p.SetState(191)
		p.Attnamelist()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserEQUALS {
		{
			p.SetState(192)
			p.Match(LuaParserEQUALS)
		}
		{
			p.SetState(193)
			p.Explist()
		}

	}

	return localctx
}

// IVariablestatContext is an interface to support dynamic dispatch.
type IVariablestatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariablestatContext differentiates from other interfaces.
	IsVariablestatContext()
}

type VariablestatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablestatContext() *VariablestatContext {
	var p = new(VariablestatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_variablestat
	return p
}

func (*VariablestatContext) IsVariablestatContext() {}

func NewVariablestatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablestatContext {
	var p = new(VariablestatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_variablestat

	return p
}

func (s *VariablestatContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablestatContext) Variablelist() IVariablelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablelistContext)
}

func (s *VariablestatContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(LuaParserEQUALS, 0)
}

func (s *VariablestatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *VariablestatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablestatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablestatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterVariablestat(s)
	}
}

func (s *VariablestatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitVariablestat(s)
	}
}

func (p *LuaParser) Variablestat() (localctx IVariablestatContext) {
	localctx = NewVariablestatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LuaParserRULE_variablestat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Variablelist()
	}
	{
		p.SetState(197)
		p.Match(LuaParserEQUALS)
	}
	{
		p.SetState(198)
		p.Explist()
	}

	return localctx
}

// IAttnamelistContext is an interface to support dynamic dispatch.
type IAttnamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttnamelistContext differentiates from other interfaces.
	IsAttnamelistContext()
}

type AttnamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttnamelistContext() *AttnamelistContext {
	var p = new(AttnamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_attnamelist
	return p
}

func (*AttnamelistContext) IsAttnamelistContext() {}

func NewAttnamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttnamelistContext {
	var p = new(AttnamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_attnamelist

	return p
}

func (s *AttnamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *AttnamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *AttnamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *AttnamelistContext) AllAttrib() []IAttribContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribContext)(nil)).Elem())
	var tst = make([]IAttribContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribContext)
		}
	}

	return tst
}

func (s *AttnamelistContext) Attrib(i int) IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *AttnamelistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LuaParserCOMMA)
}

func (s *AttnamelistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserCOMMA, i)
}

func (s *AttnamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttnamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttnamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterAttnamelist(s)
	}
}

func (s *AttnamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitAttnamelist(s)
	}
}

func (p *LuaParser) Attnamelist() (localctx IAttnamelistContext) {
	localctx = NewAttnamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LuaParserRULE_attnamelist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(201)
		p.Attrib()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserCOMMA {
		{
			p.SetState(202)
			p.Match(LuaParserCOMMA)
		}
		{
			p.SetState(203)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(204)
			p.Attrib()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAttribContext is an interface to support dynamic dispatch.
type IAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribContext differentiates from other interfaces.
	IsAttribContext()
}

type AttribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribContext() *AttribContext {
	var p = new(AttribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_attrib
	return p
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) LT() antlr.TerminalNode {
	return s.GetToken(LuaParserLT, 0)
}

func (s *AttribContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *AttribContext) GT() antlr.TerminalNode {
	return s.GetToken(LuaParserGT, 0)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (p *LuaParser) Attrib() (localctx IAttribContext) {
	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LuaParserRULE_attrib)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserLT {
		{
			p.SetState(210)
			p.Match(LuaParserLT)
		}
		{
			p.SetState(211)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(212)
			p.Match(LuaParserGT)
		}

	}

	return localctx
}

// IRetstatContext is an interface to support dynamic dispatch.
type IRetstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetstatContext differentiates from other interfaces.
	IsRetstatContext()
}

type RetstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetstatContext() *RetstatContext {
	var p = new(RetstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_retstat
	return p
}

func (*RetstatContext) IsRetstatContext() {}

func NewRetstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetstatContext {
	var p = new(RetstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_retstat

	return p
}

func (s *RetstatContext) GetParser() antlr.Parser { return s.parser }

func (s *RetstatContext) RETURN() antlr.TerminalNode {
	return s.GetToken(LuaParserRETURN, 0)
}

func (s *RetstatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *RetstatContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LuaParserSEMICOLON, 0)
}

func (s *RetstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterRetstat(s)
	}
}

func (s *RetstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitRetstat(s)
	}
}

func (p *LuaParser) Retstat() (localctx IRetstatContext) {
	localctx = NewRetstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LuaParserRULE_retstat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(LuaParserRETURN)
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserFUNCTION)|(1<<LuaParserLPAREN)|(1<<LuaParserLBRACE))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(LuaParserMINUS-39))|(1<<(LuaParserBITNOT-39))|(1<<(LuaParserNOT-39))|(1<<(LuaParserLEN-39))|(1<<(LuaParserNIL-39))|(1<<(LuaParserFALSE-39))|(1<<(LuaParserTRUE-39))|(1<<(LuaParserDOTS-39))|(1<<(LuaParserNAME-39))|(1<<(LuaParserNORMALSTRING-39))|(1<<(LuaParserCHARSTRING-39))|(1<<(LuaParserLONGSTRING-39))|(1<<(LuaParserINT-39))|(1<<(LuaParserHEX-39))|(1<<(LuaParserFLOAT-39))|(1<<(LuaParserHEX_FLOAT-39)))) != 0) {
		{
			p.SetState(216)
			p.Explist()
		}

	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserSEMICOLON {
		{
			p.SetState(219)
			p.Match(LuaParserSEMICOLON)
		}

	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) AllDCOLON() []antlr.TerminalNode {
	return s.GetTokens(LuaParserDCOLON)
}

func (s *LabelContext) DCOLON(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserDCOLON, i)
}

func (s *LabelContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *LuaParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LuaParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(LuaParserDCOLON)
	}
	{
		p.SetState(223)
		p.Match(LuaParserNAME)
	}
	{
		p.SetState(224)
		p.Match(LuaParserDCOLON)
	}

	return localctx
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *FuncnameContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *FuncnameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(LuaParserDOT)
}

func (s *FuncnameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserDOT, i)
}

func (s *FuncnameContext) COLON() antlr.TerminalNode {
	return s.GetToken(LuaParserCOLON, 0)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitFuncname(s)
	}
}

func (p *LuaParser) Funcname() (localctx IFuncnameContext) {
	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LuaParserRULE_funcname)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(LuaParserNAME)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserDOT {
		{
			p.SetState(227)
			p.Match(LuaParserDOT)
		}
		{
			p.SetState(228)
			p.Match(LuaParserNAME)
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserCOLON {
		{
			p.SetState(234)
			p.Match(LuaParserCOLON)
		}
		{
			p.SetState(235)
			p.Match(LuaParserNAME)
		}

	}

	return localctx
}

// IVariablelistContext is an interface to support dynamic dispatch.
type IVariablelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariablelistContext differentiates from other interfaces.
	IsVariablelistContext()
}

type VariablelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablelistContext() *VariablelistContext {
	var p = new(VariablelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_variablelist
	return p
}

func (*VariablelistContext) IsVariablelistContext() {}

func NewVariablelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablelistContext {
	var p = new(VariablelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_variablelist

	return p
}

func (s *VariablelistContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablelistContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *VariablelistContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariablelistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LuaParserCOMMA)
}

func (s *VariablelistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserCOMMA, i)
}

func (s *VariablelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterVariablelist(s)
	}
}

func (s *VariablelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitVariablelist(s)
	}
}

func (p *LuaParser) Variablelist() (localctx IVariablelistContext) {
	localctx = NewVariablelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LuaParserRULE_variablelist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.variable(0)
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserCOMMA {
		{
			p.SetState(239)
			p.Match(LuaParserCOMMA)
		}
		{
			p.SetState(240)
			p.variable(0)
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INamelistContext is an interface to support dynamic dispatch.
type INamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamelistContext differentiates from other interfaces.
	IsNamelistContext()
}

type NamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamelistContext() *NamelistContext {
	var p = new(NamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_namelist
	return p
}

func (*NamelistContext) IsNamelistContext() {}

func NewNamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamelistContext {
	var p = new(NamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_namelist

	return p
}

func (s *NamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *NamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(LuaParserNAME)
}

func (s *NamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, i)
}

func (s *NamelistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LuaParserCOMMA)
}

func (s *NamelistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserCOMMA, i)
}

func (s *NamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterNamelist(s)
	}
}

func (s *NamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitNamelist(s)
	}
}

func (p *LuaParser) Namelist() (localctx INamelistContext) {
	localctx = NewNamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LuaParserRULE_namelist)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(LuaParserNAME)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(247)
				p.Match(LuaParserCOMMA)
			}
			{
				p.SetState(248)
				p.Match(LuaParserNAME)
			}

		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IExplistContext is an interface to support dynamic dispatch.
type IExplistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExplistContext differentiates from other interfaces.
	IsExplistContext()
}

type ExplistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExplistContext() *ExplistContext {
	var p = new(ExplistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_explist
	return p
}

func (*ExplistContext) IsExplistContext() {}

func NewExplistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExplistContext {
	var p = new(ExplistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_explist

	return p
}

func (s *ExplistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExplistContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExplistContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExplistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LuaParserCOMMA)
}

func (s *ExplistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserCOMMA, i)
}

func (s *ExplistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExplistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterExplist(s)
	}
}

func (s *ExplistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitExplist(s)
	}
}

func (p *LuaParser) Explist() (localctx IExplistContext) {
	localctx = NewExplistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LuaParserRULE_explist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.exp(0)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserCOMMA {
		{
			p.SetState(255)
			p.Match(LuaParserCOMMA)
		}
		{
			p.SetState(256)
			p.exp(0)
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) NIL() antlr.TerminalNode {
	return s.GetToken(LuaParserNIL, 0)
}

func (s *ExpContext) FALSE() antlr.TerminalNode {
	return s.GetToken(LuaParserFALSE, 0)
}

func (s *ExpContext) TRUE() antlr.TerminalNode {
	return s.GetToken(LuaParserTRUE, 0)
}

func (s *ExpContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ExpContext) Lstring() ILstringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILstringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILstringContext)
}

func (s *ExpContext) DOTS() antlr.TerminalNode {
	return s.GetToken(LuaParserDOTS, 0)
}

func (s *ExpContext) Functiondef() IFunctiondefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctiondefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctiondefContext)
}

func (s *ExpContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ExpContext) OperatorUnary() IOperatorUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorUnaryContext)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) OperatorPower() IOperatorPowerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorPowerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorPowerContext)
}

func (s *ExpContext) OperatorMulDivMod() IOperatorMulDivModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorMulDivModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorMulDivModContext)
}

func (s *ExpContext) OperatorAddSub() IOperatorAddSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAddSubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAddSubContext)
}

func (s *ExpContext) OperatorStrcat() IOperatorStrcatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorStrcatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorStrcatContext)
}

func (s *ExpContext) OperatorComparison() IOperatorComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorComparisonContext)
}

func (s *ExpContext) OperatorAnd() IOperatorAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAndContext)
}

func (s *ExpContext) OperatorOr() IOperatorOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorOrContext)
}

func (s *ExpContext) OperatorBitwise() IOperatorBitwiseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorBitwiseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorBitwiseContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *LuaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *LuaParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, LuaParserRULE_exp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNIL:
		{
			p.SetState(263)
			p.Match(LuaParserNIL)
		}

	case LuaParserFALSE:
		{
			p.SetState(264)
			p.Match(LuaParserFALSE)
		}

	case LuaParserTRUE:
		{
			p.SetState(265)
			p.Match(LuaParserTRUE)
		}

	case LuaParserINT, LuaParserHEX, LuaParserFLOAT, LuaParserHEX_FLOAT:
		{
			p.SetState(266)
			p.Number()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		{
			p.SetState(267)
			p.Lstring()
		}

	case LuaParserDOTS:
		{
			p.SetState(268)
			p.Match(LuaParserDOTS)
		}

	case LuaParserFUNCTION:
		{
			p.SetState(269)
			p.Functiondef()
		}

	case LuaParserLPAREN, LuaParserNAME:
		{
			p.SetState(270)
			p.variable(0)
		}

	case LuaParserLBRACE:
		{
			p.SetState(271)
			p.Tableconstructor()
		}

	case LuaParserMINUS, LuaParserBITNOT, LuaParserNOT, LuaParserLEN:
		{
			p.SetState(272)
			p.OperatorUnary()
		}
		{
			p.SetState(273)
			p.exp(8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(309)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(278)
					p.OperatorPower()
				}
				{
					p.SetState(279)
					p.exp(9)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(282)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(283)
					p.exp(8)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(286)
					p.OperatorAddSub()
				}
				{
					p.SetState(287)
					p.exp(7)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(290)
					p.OperatorStrcat()
				}
				{
					p.SetState(291)
					p.exp(5)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(294)
					p.OperatorComparison()
				}
				{
					p.SetState(295)
					p.exp(5)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(298)
					p.OperatorAnd()
				}
				{
					p.SetState(299)
					p.exp(4)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(302)
					p.OperatorOr()
				}
				{
					p.SetState(303)
					p.exp(3)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_exp)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(306)
					p.OperatorBitwise()
				}
				{
					p.SetState(307)
					p.exp(2)
				}

			}

		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) CopyFrom(ctx *VariableContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NamedvariableContext struct {
	*VariableContext
}

func NewNamedvariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedvariableContext {
	var p = new(NamedvariableContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *NamedvariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedvariableContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *NamedvariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterNamedvariable(s)
	}
}

func (s *NamedvariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitNamedvariable(s)
	}
}

type ParenthesesvariableContext struct {
	*VariableContext
}

func NewParenthesesvariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesvariableContext {
	var p = new(ParenthesesvariableContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *ParenthesesvariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesvariableContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LuaParserLPAREN, 0)
}

func (s *ParenthesesvariableContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ParenthesesvariableContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LuaParserRPAREN, 0)
}

func (s *ParenthesesvariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterParenthesesvariable(s)
	}
}

func (s *ParenthesesvariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitParenthesesvariable(s)
	}
}

type FunctioncallContext struct {
	*VariableContext
}

func NewFunctioncallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctioncallContext {
	var p = new(FunctioncallContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctioncallContext) NameAndArgs() INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitFunctioncall(s)
	}
}

type IndexContext struct {
	*VariableContext
}

func NewIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexContext {
	var p = new(IndexContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *IndexContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(LuaParserLBRACK, 0)
}

func (s *IndexContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IndexContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(LuaParserRBRACK, 0)
}

func (s *IndexContext) DOT() antlr.TerminalNode {
	return s.GetToken(LuaParserDOT, 0)
}

func (s *IndexContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (p *LuaParser) Variable() (localctx IVariableContext) {
	return p.variable(0)
}

func (p *LuaParser) variable(_p int) (localctx IVariableContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewVariableContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVariableContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, LuaParserRULE_variable, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		localctx = NewNamedvariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(315)
			p.Match(LuaParserNAME)
		}

	case LuaParserLPAREN:
		localctx = NewParenthesesvariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(316)
			p.Match(LuaParserLPAREN)
		}
		{
			p.SetState(317)
			p.exp(0)
		}
		{
			p.SetState(318)
			p.Match(LuaParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(333)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewIndexContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_variable)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(329)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case LuaParserLBRACK:
					{
						p.SetState(323)
						p.Match(LuaParserLBRACK)
					}
					{
						p.SetState(324)
						p.exp(0)
					}
					{
						p.SetState(325)
						p.Match(LuaParserRBRACK)
					}

				case LuaParserDOT:
					{
						p.SetState(327)
						p.Match(LuaParserDOT)
					}
					{
						p.SetState(328)
						p.Match(LuaParserNAME)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			case 2:
				localctx = NewFunctioncallContext(p, NewVariableContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_variable)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(332)
					p.NameAndArgs()
				}

			}

		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// INameAndArgsContext is an interface to support dynamic dispatch.
type INameAndArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameAndArgsContext differentiates from other interfaces.
	IsNameAndArgsContext()
}

type NameAndArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameAndArgsContext() *NameAndArgsContext {
	var p = new(NameAndArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_nameAndArgs
	return p
}

func (*NameAndArgsContext) IsNameAndArgsContext() {}

func NewNameAndArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameAndArgsContext {
	var p = new(NameAndArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_nameAndArgs

	return p
}

func (s *NameAndArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *NameAndArgsContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *NameAndArgsContext) COLON() antlr.TerminalNode {
	return s.GetToken(LuaParserCOLON, 0)
}

func (s *NameAndArgsContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *NameAndArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameAndArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameAndArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterNameAndArgs(s)
	}
}

func (s *NameAndArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitNameAndArgs(s)
	}
}

func (p *LuaParser) NameAndArgs() (localctx INameAndArgsContext) {
	localctx = NewNameAndArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LuaParserRULE_nameAndArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserCOLON {
		{
			p.SetState(338)
			p.Match(LuaParserCOLON)
		}
		{
			p.SetState(339)
			p.Match(LuaParserNAME)
		}

	}
	{
		p.SetState(342)
		p.Args()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LuaParserLPAREN, 0)
}

func (s *ArgsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LuaParserRPAREN, 0)
}

func (s *ArgsContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *ArgsContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ArgsContext) Lstring() ILstringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILstringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILstringContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *LuaParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LuaParserRULE_args)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(351)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.Match(LuaParserLPAREN)
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserFUNCTION)|(1<<LuaParserLPAREN)|(1<<LuaParserLBRACE))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(LuaParserMINUS-39))|(1<<(LuaParserBITNOT-39))|(1<<(LuaParserNOT-39))|(1<<(LuaParserLEN-39))|(1<<(LuaParserNIL-39))|(1<<(LuaParserFALSE-39))|(1<<(LuaParserTRUE-39))|(1<<(LuaParserDOTS-39))|(1<<(LuaParserNAME-39))|(1<<(LuaParserNORMALSTRING-39))|(1<<(LuaParserCHARSTRING-39))|(1<<(LuaParserLONGSTRING-39))|(1<<(LuaParserINT-39))|(1<<(LuaParserHEX-39))|(1<<(LuaParserFLOAT-39))|(1<<(LuaParserHEX_FLOAT-39)))) != 0) {
			{
				p.SetState(345)
				p.Explist()
			}

		}
		{
			p.SetState(348)
			p.Match(LuaParserRPAREN)
		}

	case LuaParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.Tableconstructor()
		}

	case LuaParserNORMALSTRING, LuaParserCHARSTRING, LuaParserLONGSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.Lstring()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctiondefContext is an interface to support dynamic dispatch.
type IFunctiondefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctiondefContext differentiates from other interfaces.
	IsFunctiondefContext()
}

type FunctiondefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctiondefContext() *FunctiondefContext {
	var p = new(FunctiondefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_functiondef
	return p
}

func (*FunctiondefContext) IsFunctiondefContext() {}

func NewFunctiondefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctiondefContext {
	var p = new(FunctiondefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_functiondef

	return p
}

func (s *FunctiondefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctiondefContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(LuaParserFUNCTION, 0)
}

func (s *FunctiondefContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FunctiondefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiondefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctiondefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterFunctiondef(s)
	}
}

func (s *FunctiondefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitFunctiondef(s)
	}
}

func (p *LuaParser) Functiondef() (localctx IFunctiondefContext) {
	localctx = NewFunctiondefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LuaParserRULE_functiondef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(LuaParserFUNCTION)
	}
	{
		p.SetState(354)
		p.Funcbody()
	}

	return localctx
}

// IFuncbodyContext is an interface to support dynamic dispatch.
type IFuncbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncbodyContext differentiates from other interfaces.
	IsFuncbodyContext()
}

type FuncbodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncbodyContext() *FuncbodyContext {
	var p = new(FuncbodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_funcbody
	return p
}

func (*FuncbodyContext) IsFuncbodyContext() {}

func NewFuncbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncbodyContext {
	var p = new(FuncbodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_funcbody

	return p
}

func (s *FuncbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncbodyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LuaParserLPAREN, 0)
}

func (s *FuncbodyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LuaParserRPAREN, 0)
}

func (s *FuncbodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncbodyContext) END() antlr.TerminalNode {
	return s.GetToken(LuaParserEND, 0)
}

func (s *FuncbodyContext) Parlist() IParlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParlistContext)
}

func (s *FuncbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterFuncbody(s)
	}
}

func (s *FuncbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitFuncbody(s)
	}
}

func (p *LuaParser) Funcbody() (localctx IFuncbodyContext) {
	localctx = NewFuncbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LuaParserRULE_funcbody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(LuaParserLPAREN)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserDOTS || _la == LuaParserNAME {
		{
			p.SetState(357)
			p.Parlist()
		}

	}
	{
		p.SetState(360)
		p.Match(LuaParserRPAREN)
	}
	{
		p.SetState(361)
		p.Block()
	}
	{
		p.SetState(362)
		p.Match(LuaParserEND)
	}

	return localctx
}

// IParlistContext is an interface to support dynamic dispatch.
type IParlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParlistContext differentiates from other interfaces.
	IsParlistContext()
}

type ParlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParlistContext() *ParlistContext {
	var p = new(ParlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_parlist
	return p
}

func (*ParlistContext) IsParlistContext() {}

func NewParlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParlistContext {
	var p = new(ParlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_parlist

	return p
}

func (s *ParlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParlistContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *ParlistContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LuaParserCOMMA, 0)
}

func (s *ParlistContext) DOTS() antlr.TerminalNode {
	return s.GetToken(LuaParserDOTS, 0)
}

func (s *ParlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterParlist(s)
	}
}

func (s *ParlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitParlist(s)
	}
}

func (p *LuaParser) Parlist() (localctx IParlistContext) {
	localctx = NewParlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LuaParserRULE_parlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.Namelist()
		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LuaParserCOMMA {
			{
				p.SetState(365)
				p.Match(LuaParserCOMMA)
			}
			{
				p.SetState(366)
				p.Match(LuaParserDOTS)
			}

		}

	case LuaParserDOTS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(369)
			p.Match(LuaParserDOTS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableconstructorContext is an interface to support dynamic dispatch.
type ITableconstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableconstructorContext differentiates from other interfaces.
	IsTableconstructorContext()
}

type TableconstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableconstructorContext() *TableconstructorContext {
	var p = new(TableconstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_tableconstructor
	return p
}

func (*TableconstructorContext) IsTableconstructorContext() {}

func NewTableconstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableconstructorContext {
	var p = new(TableconstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_tableconstructor

	return p
}

func (s *TableconstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *TableconstructorContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LuaParserLBRACE, 0)
}

func (s *TableconstructorContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LuaParserRBRACE, 0)
}

func (s *TableconstructorContext) Fieldlist() IFieldlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldlistContext)
}

func (s *TableconstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableconstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableconstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterTableconstructor(s)
	}
}

func (s *TableconstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitTableconstructor(s)
	}
}

func (p *LuaParser) Tableconstructor() (localctx ITableconstructorContext) {
	localctx = NewTableconstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LuaParserRULE_tableconstructor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(LuaParserLBRACE)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserFUNCTION)|(1<<LuaParserLPAREN)|(1<<LuaParserLBRACK)|(1<<LuaParserLBRACE))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(LuaParserMINUS-39))|(1<<(LuaParserBITNOT-39))|(1<<(LuaParserNOT-39))|(1<<(LuaParserLEN-39))|(1<<(LuaParserNIL-39))|(1<<(LuaParserFALSE-39))|(1<<(LuaParserTRUE-39))|(1<<(LuaParserDOTS-39))|(1<<(LuaParserNAME-39))|(1<<(LuaParserNORMALSTRING-39))|(1<<(LuaParserCHARSTRING-39))|(1<<(LuaParserLONGSTRING-39))|(1<<(LuaParserINT-39))|(1<<(LuaParserHEX-39))|(1<<(LuaParserFLOAT-39))|(1<<(LuaParserHEX_FLOAT-39)))) != 0) {
		{
			p.SetState(373)
			p.Fieldlist()
		}

	}
	{
		p.SetState(376)
		p.Match(LuaParserRBRACE)
	}

	return localctx
}

// IFieldlistContext is an interface to support dynamic dispatch.
type IFieldlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldlistContext differentiates from other interfaces.
	IsFieldlistContext()
}

type FieldlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldlistContext() *FieldlistContext {
	var p = new(FieldlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_fieldlist
	return p
}

func (*FieldlistContext) IsFieldlistContext() {}

func NewFieldlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldlistContext {
	var p = new(FieldlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_fieldlist

	return p
}

func (s *FieldlistContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldlistContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldlistContext) AllFieldsep() []IFieldsepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldsepContext)(nil)).Elem())
	var tst = make([]IFieldsepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldsepContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Fieldsep(i int) IFieldsepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldsepContext)
}

func (s *FieldlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterFieldlist(s)
	}
}

func (s *FieldlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitFieldlist(s)
	}
}

func (p *LuaParser) Fieldlist() (localctx IFieldlistContext) {
	localctx = NewFieldlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LuaParserRULE_fieldlist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Field()
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(379)
				p.Fieldsep()
			}
			{
				p.SetState(380)
				p.Field()
			}

		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserSEMICOLON || _la == LuaParserCOMMA {
		{
			p.SetState(387)
			p.Fieldsep()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(LuaParserLBRACK, 0)
}

func (s *FieldContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *FieldContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FieldContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(LuaParserRBRACK, 0)
}

func (s *FieldContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(LuaParserEQUALS, 0)
}

func (s *FieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(LuaParserNAME, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *LuaParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LuaParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.Match(LuaParserLBRACK)
		}
		{
			p.SetState(391)
			p.exp(0)
		}
		{
			p.SetState(392)
			p.Match(LuaParserRBRACK)
		}
		{
			p.SetState(393)
			p.Match(LuaParserEQUALS)
		}
		{
			p.SetState(394)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Match(LuaParserNAME)
		}
		{
			p.SetState(397)
			p.Match(LuaParserEQUALS)
		}
		{
			p.SetState(398)
			p.exp(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(399)
			p.exp(0)
		}

	}

	return localctx
}

// IFieldsepContext is an interface to support dynamic dispatch.
type IFieldsepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsepContext differentiates from other interfaces.
	IsFieldsepContext()
}

type FieldsepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsepContext() *FieldsepContext {
	var p = new(FieldsepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_fieldsep
	return p
}

func (*FieldsepContext) IsFieldsepContext() {}

func NewFieldsepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsepContext {
	var p = new(FieldsepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_fieldsep

	return p
}

func (s *FieldsepContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsepContext) COMMA() antlr.TerminalNode {
	return s.GetToken(LuaParserCOMMA, 0)
}

func (s *FieldsepContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(LuaParserSEMICOLON, 0)
}

func (s *FieldsepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterFieldsep(s)
	}
}

func (s *FieldsepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitFieldsep(s)
	}
}

func (p *LuaParser) Fieldsep() (localctx IFieldsepContext) {
	localctx = NewFieldsepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LuaParserRULE_fieldsep)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserSEMICOLON || _la == LuaParserCOMMA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorOrContext is an interface to support dynamic dispatch.
type IOperatorOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorOrContext differentiates from other interfaces.
	IsOperatorOrContext()
}

type OperatorOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorOrContext() *OperatorOrContext {
	var p = new(OperatorOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorOr
	return p
}

func (*OperatorOrContext) IsOperatorOrContext() {}

func NewOperatorOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorOrContext {
	var p = new(OperatorOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorOr

	return p
}

func (s *OperatorOrContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorOrContext) OR() antlr.TerminalNode {
	return s.GetToken(LuaParserOR, 0)
}

func (s *OperatorOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorOr(s)
	}
}

func (s *OperatorOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorOr(s)
	}
}

func (p *LuaParser) OperatorOr() (localctx IOperatorOrContext) {
	localctx = NewOperatorOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LuaParserRULE_operatorOr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Match(LuaParserOR)
	}

	return localctx
}

// IOperatorAndContext is an interface to support dynamic dispatch.
type IOperatorAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAndContext differentiates from other interfaces.
	IsOperatorAndContext()
}

type OperatorAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAndContext() *OperatorAndContext {
	var p = new(OperatorAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAnd
	return p
}

func (*OperatorAndContext) IsOperatorAndContext() {}

func NewOperatorAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAndContext {
	var p = new(OperatorAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorAnd

	return p
}

func (s *OperatorAndContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorAndContext) AND() antlr.TerminalNode {
	return s.GetToken(LuaParserAND, 0)
}

func (s *OperatorAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorAnd(s)
	}
}

func (s *OperatorAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorAnd(s)
	}
}

func (p *LuaParser) OperatorAnd() (localctx IOperatorAndContext) {
	localctx = NewOperatorAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LuaParserRULE_operatorAnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(LuaParserAND)
	}

	return localctx
}

// IOperatorComparisonContext is an interface to support dynamic dispatch.
type IOperatorComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorComparisonContext differentiates from other interfaces.
	IsOperatorComparisonContext()
}

type OperatorComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorComparisonContext() *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorComparison
	return p
}

func (*OperatorComparisonContext) IsOperatorComparisonContext() {}

func NewOperatorComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorComparison

	return p
}

func (s *OperatorComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorComparisonContext) CopyFrom(ctx *OperatorComparisonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperatorComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualContext struct {
	*OperatorComparisonContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	p.OperatorComparisonContext = NewEmptyOperatorComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorComparisonContext))

	return p
}

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) EQ() antlr.TerminalNode {
	return s.GetToken(LuaParserEQ, 0)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitEqual(s)
	}
}

type LessthanContext struct {
	*OperatorComparisonContext
}

func NewLessthanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessthanContext {
	var p = new(LessthanContext)

	p.OperatorComparisonContext = NewEmptyOperatorComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorComparisonContext))

	return p
}

func (s *LessthanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessthanContext) LT() antlr.TerminalNode {
	return s.GetToken(LuaParserLT, 0)
}

func (s *LessthanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterLessthan(s)
	}
}

func (s *LessthanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitLessthan(s)
	}
}

type LessthanorequalContext struct {
	*OperatorComparisonContext
}

func NewLessthanorequalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessthanorequalContext {
	var p = new(LessthanorequalContext)

	p.OperatorComparisonContext = NewEmptyOperatorComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorComparisonContext))

	return p
}

func (s *LessthanorequalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessthanorequalContext) LTE() antlr.TerminalNode {
	return s.GetToken(LuaParserLTE, 0)
}

func (s *LessthanorequalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterLessthanorequal(s)
	}
}

func (s *LessthanorequalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitLessthanorequal(s)
	}
}

type GreaterthanContext struct {
	*OperatorComparisonContext
}

func NewGreaterthanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterthanContext {
	var p = new(GreaterthanContext)

	p.OperatorComparisonContext = NewEmptyOperatorComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorComparisonContext))

	return p
}

func (s *GreaterthanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterthanContext) GT() antlr.TerminalNode {
	return s.GetToken(LuaParserGT, 0)
}

func (s *GreaterthanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterGreaterthan(s)
	}
}

func (s *GreaterthanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitGreaterthan(s)
	}
}

type NotequalContext struct {
	*OperatorComparisonContext
}

func NewNotequalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotequalContext {
	var p = new(NotequalContext)

	p.OperatorComparisonContext = NewEmptyOperatorComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorComparisonContext))

	return p
}

func (s *NotequalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotequalContext) NEQ() antlr.TerminalNode {
	return s.GetToken(LuaParserNEQ, 0)
}

func (s *NotequalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterNotequal(s)
	}
}

func (s *NotequalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitNotequal(s)
	}
}

type GreaterthanorequalContext struct {
	*OperatorComparisonContext
}

func NewGreaterthanorequalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterthanorequalContext {
	var p = new(GreaterthanorequalContext)

	p.OperatorComparisonContext = NewEmptyOperatorComparisonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorComparisonContext))

	return p
}

func (s *GreaterthanorequalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterthanorequalContext) GTE() antlr.TerminalNode {
	return s.GetToken(LuaParserGTE, 0)
}

func (s *GreaterthanorequalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterGreaterthanorequal(s)
	}
}

func (s *GreaterthanorequalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitGreaterthanorequal(s)
	}
}

func (p *LuaParser) OperatorComparison() (localctx IOperatorComparisonContext) {
	localctx = NewOperatorComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, LuaParserRULE_operatorComparison)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(414)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserLT:
		localctx = NewLessthanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(408)
			p.Match(LuaParserLT)
		}

	case LuaParserGT:
		localctx = NewGreaterthanContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(409)
			p.Match(LuaParserGT)
		}

	case LuaParserLTE:
		localctx = NewLessthanorequalContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(410)
			p.Match(LuaParserLTE)
		}

	case LuaParserGTE:
		localctx = NewGreaterthanorequalContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(411)
			p.Match(LuaParserGTE)
		}

	case LuaParserNEQ:
		localctx = NewNotequalContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(412)
			p.Match(LuaParserNEQ)
		}

	case LuaParserEQ:
		localctx = NewEqualContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(413)
			p.Match(LuaParserEQ)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperatorStrcatContext is an interface to support dynamic dispatch.
type IOperatorStrcatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorStrcatContext differentiates from other interfaces.
	IsOperatorStrcatContext()
}

type OperatorStrcatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorStrcatContext() *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorStrcat
	return p
}

func (*OperatorStrcatContext) IsOperatorStrcatContext() {}

func NewOperatorStrcatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorStrcat

	return p
}

func (s *OperatorStrcatContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorStrcatContext) STRCAT() antlr.TerminalNode {
	return s.GetToken(LuaParserSTRCAT, 0)
}

func (s *OperatorStrcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorStrcatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorStrcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorStrcat(s)
	}
}

func (s *OperatorStrcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorStrcat(s)
	}
}

func (p *LuaParser) OperatorStrcat() (localctx IOperatorStrcatContext) {
	localctx = NewOperatorStrcatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, LuaParserRULE_operatorStrcat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.Match(LuaParserSTRCAT)
	}

	return localctx
}

// IOperatorAddSubContext is an interface to support dynamic dispatch.
type IOperatorAddSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAddSubContext differentiates from other interfaces.
	IsOperatorAddSubContext()
}

type OperatorAddSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAddSubContext() *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorAddSub
	return p
}

func (*OperatorAddSubContext) IsOperatorAddSubContext() {}

func NewOperatorAddSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorAddSub

	return p
}

func (s *OperatorAddSubContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorAddSubContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LuaParserPLUS, 0)
}

func (s *OperatorAddSubContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LuaParserMINUS, 0)
}

func (s *OperatorAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAddSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorAddSub(s)
	}
}

func (s *OperatorAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorAddSub(s)
	}
}

func (p *LuaParser) OperatorAddSub() (localctx IOperatorAddSubContext) {
	localctx = NewOperatorAddSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, LuaParserRULE_operatorAddSub)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserPLUS || _la == LuaParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorMulDivModContext is an interface to support dynamic dispatch.
type IOperatorMulDivModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorMulDivModContext differentiates from other interfaces.
	IsOperatorMulDivModContext()
}

type OperatorMulDivModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorMulDivModContext() *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorMulDivMod
	return p
}

func (*OperatorMulDivModContext) IsOperatorMulDivModContext() {}

func NewOperatorMulDivModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorMulDivMod

	return p
}

func (s *OperatorMulDivModContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorMulDivModContext) MUL() antlr.TerminalNode {
	return s.GetToken(LuaParserMUL, 0)
}

func (s *OperatorMulDivModContext) DIV() antlr.TerminalNode {
	return s.GetToken(LuaParserDIV, 0)
}

func (s *OperatorMulDivModContext) MOD() antlr.TerminalNode {
	return s.GetToken(LuaParserMOD, 0)
}

func (s *OperatorMulDivModContext) DIVFLOOR() antlr.TerminalNode {
	return s.GetToken(LuaParserDIVFLOOR, 0)
}

func (s *OperatorMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorMulDivModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorMulDivMod(s)
	}
}

func (s *OperatorMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorMulDivMod(s)
	}
}

func (p *LuaParser) OperatorMulDivMod() (localctx IOperatorMulDivModContext) {
	localctx = NewOperatorMulDivModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, LuaParserRULE_operatorMulDivMod)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(LuaParserMUL-40))|(1<<(LuaParserDIV-40))|(1<<(LuaParserMOD-40))|(1<<(LuaParserDIVFLOOR-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorBitwiseContext is an interface to support dynamic dispatch.
type IOperatorBitwiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorBitwiseContext differentiates from other interfaces.
	IsOperatorBitwiseContext()
}

type OperatorBitwiseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorBitwiseContext() *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorBitwise
	return p
}

func (*OperatorBitwiseContext) IsOperatorBitwiseContext() {}

func NewOperatorBitwiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorBitwise

	return p
}

func (s *OperatorBitwiseContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorBitwiseContext) BITAND() antlr.TerminalNode {
	return s.GetToken(LuaParserBITAND, 0)
}

func (s *OperatorBitwiseContext) BITOR() antlr.TerminalNode {
	return s.GetToken(LuaParserBITOR, 0)
}

func (s *OperatorBitwiseContext) BITNOT() antlr.TerminalNode {
	return s.GetToken(LuaParserBITNOT, 0)
}

func (s *OperatorBitwiseContext) BITSHL() antlr.TerminalNode {
	return s.GetToken(LuaParserBITSHL, 0)
}

func (s *OperatorBitwiseContext) BITSHR() antlr.TerminalNode {
	return s.GetToken(LuaParserBITSHR, 0)
}

func (s *OperatorBitwiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorBitwiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorBitwiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorBitwise(s)
	}
}

func (s *OperatorBitwiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorBitwise(s)
	}
}

func (p *LuaParser) OperatorBitwise() (localctx IOperatorBitwiseContext) {
	localctx = NewOperatorBitwiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, LuaParserRULE_operatorBitwise)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(LuaParserBITAND-44))|(1<<(LuaParserBITOR-44))|(1<<(LuaParserBITNOT-44))|(1<<(LuaParserBITSHL-44))|(1<<(LuaParserBITSHR-44)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorUnaryContext is an interface to support dynamic dispatch.
type IOperatorUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorUnaryContext differentiates from other interfaces.
	IsOperatorUnaryContext()
}

type OperatorUnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorUnaryContext() *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorUnary
	return p
}

func (*OperatorUnaryContext) IsOperatorUnaryContext() {}

func NewOperatorUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorUnary

	return p
}

func (s *OperatorUnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorUnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(LuaParserNOT, 0)
}

func (s *OperatorUnaryContext) LEN() antlr.TerminalNode {
	return s.GetToken(LuaParserLEN, 0)
}

func (s *OperatorUnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LuaParserMINUS, 0)
}

func (s *OperatorUnaryContext) BITNOT() antlr.TerminalNode {
	return s.GetToken(LuaParserBITNOT, 0)
}

func (s *OperatorUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorUnary(s)
	}
}

func (s *OperatorUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorUnary(s)
	}
}

func (p *LuaParser) OperatorUnary() (localctx IOperatorUnaryContext) {
	localctx = NewOperatorUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, LuaParserRULE_operatorUnary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(LuaParserMINUS-39))|(1<<(LuaParserBITNOT-39))|(1<<(LuaParserNOT-39))|(1<<(LuaParserLEN-39)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorPowerContext is an interface to support dynamic dispatch.
type IOperatorPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorPowerContext differentiates from other interfaces.
	IsOperatorPowerContext()
}

type OperatorPowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorPowerContext() *OperatorPowerContext {
	var p = new(OperatorPowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_operatorPower
	return p
}

func (*OperatorPowerContext) IsOperatorPowerContext() {}

func NewOperatorPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorPowerContext {
	var p = new(OperatorPowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_operatorPower

	return p
}

func (s *OperatorPowerContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorPowerContext) POWER() antlr.TerminalNode {
	return s.GetToken(LuaParserPOWER, 0)
}

func (s *OperatorPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorPowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorPowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterOperatorPower(s)
	}
}

func (s *OperatorPowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitOperatorPower(s)
	}
}

func (p *LuaParser) OperatorPower() (localctx IOperatorPowerContext) {
	localctx = NewOperatorPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, LuaParserRULE_operatorPower)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(LuaParserPOWER)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(LuaParserINT, 0)
}

func (s *NumberContext) HEX() antlr.TerminalNode {
	return s.GetToken(LuaParserHEX, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(LuaParserFLOAT, 0)
}

func (s *NumberContext) HEX_FLOAT() antlr.TerminalNode {
	return s.GetToken(LuaParserHEX_FLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *LuaParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, LuaParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(428)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(LuaParserINT-60))|(1<<(LuaParserHEX-60))|(1<<(LuaParserFLOAT-60))|(1<<(LuaParserHEX_FLOAT-60)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILstringContext is an interface to support dynamic dispatch.
type ILstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLstringContext differentiates from other interfaces.
	IsLstringContext()
}

type LstringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLstringContext() *LstringContext {
	var p = new(LstringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_lstring
	return p
}

func (*LstringContext) IsLstringContext() {}

func NewLstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LstringContext {
	var p = new(LstringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_lstring

	return p
}

func (s *LstringContext) GetParser() antlr.Parser { return s.parser }

func (s *LstringContext) NORMALSTRING() antlr.TerminalNode {
	return s.GetToken(LuaParserNORMALSTRING, 0)
}

func (s *LstringContext) CHARSTRING() antlr.TerminalNode {
	return s.GetToken(LuaParserCHARSTRING, 0)
}

func (s *LstringContext) LONGSTRING() antlr.TerminalNode {
	return s.GetToken(LuaParserLONGSTRING, 0)
}

func (s *LstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LstringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LstringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.EnterLstring(s)
	}
}

func (s *LstringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LuaParserListener); ok {
		listenerT.ExitLstring(s)
	}
}

func (p *LuaParser) Lstring() (localctx ILstringContext) {
	localctx = NewLstringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, LuaParserRULE_lstring)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(LuaParserNORMALSTRING-57))|(1<<(LuaParserCHARSTRING-57))|(1<<(LuaParserLONGSTRING-57)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *LuaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 23:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	case 24:
		var t *VariableContext = nil
		if localctx != nil {
			t = localctx.(*VariableContext)
		}
		return p.Variable_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LuaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LuaParser) Variable_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
