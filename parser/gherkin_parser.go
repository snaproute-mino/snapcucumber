// Generated from Gherkin.g4 by ANTLR 4.7.

package parser // Gherkin

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 364,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 6, 2, 44, 10, 2, 13,
	2, 14, 2, 45, 3, 3, 3, 3, 3, 3, 3, 4, 6, 4, 52, 10, 4, 13, 4, 14, 4, 53,
	3, 5, 5, 5, 57, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 65, 10,
	6, 3, 6, 3, 6, 3, 6, 6, 6, 70, 10, 6, 13, 6, 14, 6, 71, 5, 6, 74, 10, 6,
	3, 6, 7, 6, 77, 10, 6, 12, 6, 14, 6, 80, 11, 6, 3, 6, 7, 6, 83, 10, 6,
	12, 6, 14, 6, 86, 11, 6, 3, 6, 3, 6, 3, 6, 5, 6, 91, 10, 6, 3, 6, 5, 6,
	94, 10, 6, 3, 7, 3, 7, 3, 8, 5, 8, 99, 10, 8, 3, 8, 3, 8, 7, 8, 103, 10,
	8, 12, 8, 14, 8, 106, 11, 8, 3, 8, 3, 8, 5, 8, 110, 10, 8, 3, 8, 5, 8,
	113, 10, 8, 3, 9, 6, 9, 116, 10, 9, 13, 9, 14, 9, 117, 3, 10, 7, 10, 121,
	10, 10, 12, 10, 14, 10, 124, 11, 10, 3, 10, 5, 10, 127, 10, 10, 3, 10,
	3, 10, 5, 10, 131, 10, 10, 3, 10, 3, 10, 5, 10, 135, 10, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 140, 10, 10, 3, 10, 5, 10, 143, 10, 10, 3, 10, 7, 10,
	146, 10, 10, 12, 10, 14, 10, 149, 11, 10, 3, 10, 5, 10, 152, 10, 10, 3,
	10, 3, 10, 7, 10, 156, 10, 10, 12, 10, 14, 10, 159, 11, 10, 3, 10, 3, 10,
	3, 11, 7, 11, 164, 10, 11, 12, 11, 14, 11, 167, 11, 11, 3, 11, 5, 11, 170,
	10, 11, 3, 11, 3, 11, 5, 11, 174, 10, 11, 3, 11, 3, 11, 5, 11, 178, 10,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 184, 10, 11, 3, 11, 6, 11, 187,
	10, 11, 13, 11, 14, 11, 188, 3, 12, 3, 12, 5, 12, 193, 10, 12, 3, 13, 7,
	13, 196, 10, 13, 12, 13, 14, 13, 199, 11, 13, 3, 13, 5, 13, 202, 10, 13,
	3, 13, 3, 13, 5, 13, 206, 10, 13, 3, 13, 3, 13, 5, 13, 210, 10, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 216, 10, 13, 3, 13, 7, 13, 219, 10, 13,
	12, 13, 14, 13, 222, 11, 13, 3, 14, 7, 14, 225, 10, 14, 12, 14, 14, 14,
	228, 11, 14, 3, 14, 5, 14, 231, 10, 14, 3, 14, 3, 14, 5, 14, 235, 10, 14,
	3, 14, 3, 14, 5, 14, 239, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 245,
	10, 14, 3, 14, 7, 14, 248, 10, 14, 12, 14, 14, 14, 251, 11, 14, 3, 14,
	6, 14, 254, 10, 14, 13, 14, 14, 14, 255, 3, 14, 5, 14, 259, 10, 14, 3,
	14, 7, 14, 262, 10, 14, 12, 14, 14, 14, 265, 11, 14, 3, 15, 7, 15, 268,
	10, 15, 12, 15, 14, 15, 271, 11, 15, 3, 15, 5, 15, 274, 10, 15, 3, 15,
	3, 15, 5, 15, 278, 10, 15, 3, 15, 3, 15, 5, 15, 282, 10, 15, 3, 15, 5,
	15, 285, 10, 15, 3, 15, 3, 15, 6, 15, 289, 10, 15, 13, 15, 14, 15, 290,
	3, 16, 7, 16, 294, 10, 16, 12, 16, 14, 16, 297, 11, 16, 3, 16, 5, 16, 300,
	10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 6, 16, 307, 10, 16, 13, 16,
	14, 16, 308, 3, 16, 6, 16, 312, 10, 16, 13, 16, 14, 16, 313, 5, 16, 316,
	10, 16, 3, 17, 7, 17, 319, 10, 17, 12, 17, 14, 17, 322, 11, 17, 3, 17,
	5, 17, 325, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 331, 10, 17, 13,
	17, 14, 17, 332, 3, 17, 5, 17, 336, 10, 17, 3, 17, 3, 17, 3, 18, 7, 18,
	341, 10, 18, 12, 18, 14, 18, 344, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 5, 19, 351, 10, 19, 3, 19, 3, 19, 3, 20, 5, 20, 356, 10, 20, 3, 21,
	7, 21, 359, 10, 21, 12, 21, 14, 21, 362, 11, 21, 3, 21, 3, 360, 2, 22,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 2, 7, 3, 2, 9, 10, 3, 2, 9, 9, 4, 2, 3, 11, 14, 15, 4, 2, 9, 9, 11,
	11, 3, 2, 16, 16, 2, 414, 2, 43, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 51,
	3, 2, 2, 2, 8, 56, 3, 2, 2, 2, 10, 93, 3, 2, 2, 2, 12, 95, 3, 2, 2, 2,
	14, 112, 3, 2, 2, 2, 16, 115, 3, 2, 2, 2, 18, 122, 3, 2, 2, 2, 20, 165,
	3, 2, 2, 2, 22, 192, 3, 2, 2, 2, 24, 197, 3, 2, 2, 2, 26, 226, 3, 2, 2,
	2, 28, 269, 3, 2, 2, 2, 30, 295, 3, 2, 2, 2, 32, 320, 3, 2, 2, 2, 34, 342,
	3, 2, 2, 2, 36, 345, 3, 2, 2, 2, 38, 355, 3, 2, 2, 2, 40, 360, 3, 2, 2,
	2, 42, 44, 10, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 43,
	3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 3, 3, 2, 2, 2, 47, 48, 7, 14, 2, 2,
	48, 49, 5, 2, 2, 2, 49, 5, 3, 2, 2, 2, 50, 52, 10, 3, 2, 2, 51, 50, 3,
	2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54,
	7, 3, 2, 2, 2, 55, 57, 7, 10, 2, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2,
	2, 57, 58, 3, 2, 2, 2, 58, 59, 7, 15, 2, 2, 59, 60, 5, 6, 4, 2, 60, 61,
	7, 9, 2, 2, 61, 9, 3, 2, 2, 2, 62, 94, 5, 8, 5, 2, 63, 65, 7, 10, 2, 2,
	64, 63, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 78, 5,
	4, 3, 2, 67, 74, 7, 10, 2, 2, 68, 70, 7, 9, 2, 2, 69, 68, 3, 2, 2, 2, 70,
	71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2,
	2, 73, 67, 3, 2, 2, 2, 73, 69, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 77,
	5, 4, 3, 2, 76, 73, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2,
	78, 79, 3, 2, 2, 2, 79, 84, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 83, 7,
	10, 2, 2, 82, 81, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 87, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 88, 7, 9, 2,
	2, 88, 94, 3, 2, 2, 2, 89, 91, 7, 10, 2, 2, 90, 89, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94, 7, 9, 2, 2, 93, 62, 3, 2, 2, 2,
	93, 64, 3, 2, 2, 2, 93, 90, 3, 2, 2, 2, 94, 11, 3, 2, 2, 2, 95, 96, 5,
	8, 5, 2, 96, 13, 3, 2, 2, 2, 97, 99, 7, 10, 2, 2, 98, 97, 3, 2, 2, 2, 98,
	99, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 104, 10, 4, 2, 2, 101, 103, 10,
	3, 2, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2,
	2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107,
	113, 7, 9, 2, 2, 108, 110, 7, 10, 2, 2, 109, 108, 3, 2, 2, 2, 109, 110,
	3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 113, 7, 9, 2, 2, 112, 98, 3, 2,
	2, 2, 112, 109, 3, 2, 2, 2, 113, 15, 3, 2, 2, 2, 114, 116, 5, 14, 8, 2,
	115, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117,
	118, 3, 2, 2, 2, 118, 17, 3, 2, 2, 2, 119, 121, 5, 10, 6, 2, 120, 119,
	3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2,
	2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 127, 7, 10, 2, 2,
	126, 125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128,
	130, 7, 4, 2, 2, 129, 131, 7, 10, 2, 2, 130, 129, 3, 2, 2, 2, 130, 131,
	3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 7, 13, 2, 2, 133, 135, 7, 10,
	2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2,
	136, 137, 5, 6, 4, 2, 137, 139, 7, 9, 2, 2, 138, 140, 5, 16, 9, 2, 139,
	138, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142, 3, 2, 2, 2, 141, 143,
	5, 20, 11, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 147, 3,
	2, 2, 2, 144, 146, 5, 22, 12, 2, 145, 144, 3, 2, 2, 2, 146, 149, 3, 2,
	2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 157, 3, 2, 2, 2,
	149, 147, 3, 2, 2, 2, 150, 152, 7, 10, 2, 2, 151, 150, 3, 2, 2, 2, 151,
	152, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 156, 7, 9, 2, 2, 154, 156,
	5, 12, 7, 2, 155, 151, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156, 159, 3, 2,
	2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 160, 3, 2, 2, 2,
	159, 157, 3, 2, 2, 2, 160, 161, 7, 2, 2, 3, 161, 19, 3, 2, 2, 2, 162, 164,
	5, 10, 6, 2, 163, 162, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2,
	2, 2, 165, 166, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2,
	168, 170, 7, 10, 2, 2, 169, 168, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170,
	171, 3, 2, 2, 2, 171, 173, 7, 5, 2, 2, 172, 174, 7, 10, 2, 2, 173, 172,
	3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 7, 13,
	2, 2, 176, 178, 7, 10, 2, 2, 177, 176, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2,
	178, 179, 3, 2, 2, 2, 179, 180, 5, 6, 4, 2, 180, 186, 7, 9, 2, 2, 181,
	187, 5, 30, 16, 2, 182, 184, 7, 10, 2, 2, 183, 182, 3, 2, 2, 2, 183, 184,
	3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 7, 9, 2, 2, 186, 181, 3, 2,
	2, 2, 186, 183, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2,
	188, 189, 3, 2, 2, 2, 189, 21, 3, 2, 2, 2, 190, 193, 5, 24, 13, 2, 191,
	193, 5, 26, 14, 2, 192, 190, 3, 2, 2, 2, 192, 191, 3, 2, 2, 2, 193, 23,
	3, 2, 2, 2, 194, 196, 5, 10, 6, 2, 195, 194, 3, 2, 2, 2, 196, 199, 3, 2,
	2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2,
	199, 197, 3, 2, 2, 2, 200, 202, 7, 10, 2, 2, 201, 200, 3, 2, 2, 2, 201,
	202, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205, 7, 6, 2, 2, 204, 206,
	7, 10, 2, 2, 205, 204, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3, 2,
	2, 2, 207, 209, 7, 13, 2, 2, 208, 210, 7, 10, 2, 2, 209, 208, 3, 2, 2,
	2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 5, 6, 4, 2, 212,
	220, 7, 9, 2, 2, 213, 219, 5, 30, 16, 2, 214, 216, 7, 10, 2, 2, 215, 214,
	3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 219, 7, 9,
	2, 2, 218, 213, 3, 2, 2, 2, 218, 215, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2,
	220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 25, 3, 2, 2, 2, 222, 220,
	3, 2, 2, 2, 223, 225, 5, 10, 6, 2, 224, 223, 3, 2, 2, 2, 225, 228, 3, 2,
	2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2,
	228, 226, 3, 2, 2, 2, 229, 231, 7, 10, 2, 2, 230, 229, 3, 2, 2, 2, 230,
	231, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 234, 7, 7, 2, 2, 233, 235,
	7, 10, 2, 2, 234, 233, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 3, 2,
	2, 2, 236, 238, 7, 13, 2, 2, 237, 239, 7, 10, 2, 2, 238, 237, 3, 2, 2,
	2, 238, 239, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 241, 5, 6, 4, 2, 241,
	249, 7, 9, 2, 2, 242, 248, 5, 30, 16, 2, 243, 245, 7, 10, 2, 2, 244, 243,
	3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 248, 7, 9,
	2, 2, 247, 242, 3, 2, 2, 2, 247, 244, 3, 2, 2, 2, 248, 251, 3, 2, 2, 2,
	249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 253, 3, 2, 2, 2, 251,
	249, 3, 2, 2, 2, 252, 254, 5, 28, 15, 2, 253, 252, 3, 2, 2, 2, 254, 255,
	3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 263, 3, 2,
	2, 2, 257, 259, 7, 10, 2, 2, 258, 257, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2,
	259, 260, 3, 2, 2, 2, 260, 262, 7, 9, 2, 2, 261, 258, 3, 2, 2, 2, 262,
	265, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 27, 3,
	2, 2, 2, 265, 263, 3, 2, 2, 2, 266, 268, 5, 10, 6, 2, 267, 266, 3, 2, 2,
	2, 268, 271, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270,
	273, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272, 274, 7, 10, 2, 2, 273, 272,
	3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 277, 7, 8,
	2, 2, 276, 278, 7, 10, 2, 2, 277, 276, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2,
	278, 279, 3, 2, 2, 2, 279, 281, 7, 13, 2, 2, 280, 282, 7, 10, 2, 2, 281,
	280, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 284, 3, 2, 2, 2, 283, 285,
	5, 6, 4, 2, 284, 283, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 286, 3, 2,
	2, 2, 286, 288, 7, 9, 2, 2, 287, 289, 5, 32, 17, 2, 288, 287, 3, 2, 2,
	2, 289, 290, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291,
	29, 3, 2, 2, 2, 292, 294, 5, 10, 6, 2, 293, 292, 3, 2, 2, 2, 294, 297,
	3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 299, 3, 2,
	2, 2, 297, 295, 3, 2, 2, 2, 298, 300, 7, 10, 2, 2, 299, 298, 3, 2, 2, 2,
	299, 300, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 302, 7, 3, 2, 2, 302,
	303, 7, 10, 2, 2, 303, 304, 5, 6, 4, 2, 304, 315, 7, 9, 2, 2, 305, 307,
	5, 32, 17, 2, 306, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 306, 3,
	2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 316, 3, 2, 2, 2, 310, 312, 5, 36, 19,
	2, 311, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313,
	314, 3, 2, 2, 2, 314, 316, 3, 2, 2, 2, 315, 306, 3, 2, 2, 2, 315, 311,
	3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 31, 3, 2, 2, 2, 317, 319, 5, 10,
	6, 2, 318, 317, 3, 2, 2, 2, 319, 322, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2,
	320, 321, 3, 2, 2, 2, 321, 324, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 323,
	325, 7, 10, 2, 2, 324, 323, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 326,
	3, 2, 2, 2, 326, 330, 7, 11, 2, 2, 327, 328, 5, 34, 18, 2, 328, 329, 7,
	11, 2, 2, 329, 331, 3, 2, 2, 2, 330, 327, 3, 2, 2, 2, 331, 332, 3, 2, 2,
	2, 332, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 335, 3, 2, 2, 2, 334,
	336, 7, 10, 2, 2, 335, 334, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 337,
	3, 2, 2, 2, 337, 338, 7, 9, 2, 2, 338, 33, 3, 2, 2, 2, 339, 341, 10, 5,
	2, 2, 340, 339, 3, 2, 2, 2, 341, 344, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2,
	342, 343, 3, 2, 2, 2, 343, 35, 3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 345, 346,
	5, 38, 20, 2, 346, 347, 7, 16, 2, 2, 347, 348, 5, 40, 21, 2, 348, 350,
	7, 16, 2, 2, 349, 351, 7, 10, 2, 2, 350, 349, 3, 2, 2, 2, 350, 351, 3,
	2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 353, 7, 9, 2, 2, 353, 37, 3, 2, 2,
	2, 354, 356, 7, 10, 2, 2, 355, 354, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356,
	39, 3, 2, 2, 2, 357, 359, 10, 6, 2, 2, 358, 357, 3, 2, 2, 2, 359, 362,
	3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 361, 41, 3, 2,
	2, 2, 362, 360, 3, 2, 2, 2, 71, 45, 53, 56, 64, 71, 73, 78, 84, 90, 93,
	98, 104, 109, 112, 117, 122, 126, 130, 134, 139, 142, 147, 151, 155, 157,
	165, 169, 173, 177, 183, 186, 188, 192, 197, 201, 205, 209, 215, 218, 220,
	226, 230, 234, 238, 244, 247, 249, 255, 258, 263, 269, 273, 277, 281, 284,
	290, 295, 299, 308, 313, 315, 320, 324, 332, 335, 342, 350, 355, 360,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'Feature'", "'Background'", "'Scenario'", "'Scenario Outline'",
	"'Examples'", "", "", "'|'", "'\\|'", "':'", "'@'", "'#'", "'\"\"\"'",
	"'\"'",
}
var symbolicNames = []string{
	"", "STEP_KW", "FEATURE_KW", "BACKGROUND_KW", "SCENARIO_KW", "OUTLINE_KW",
	"EXAMPLES_KW", "EOL", "WS", "PIPE", "ESCAPED_PIPE", "COLON", "AT", "HASH",
	"TRIPLE_QUOTE", "QUOTE", "CHAR",
}

var ruleNames = []string{
	"tagName", "tag", "lineContent", "comment", "annotation", "trailingComment",
	"descriptionLine", "name", "feature", "background", "abstractScenario",
	"scenario", "outline", "examples", "step", "row", "cell", "document", "documentIndent",
	"documentContent",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GherkinParser struct {
	*antlr.BaseParser
}

func NewGherkinParser(input antlr.TokenStream) *GherkinParser {
	this := new(GherkinParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gherkin.g4"

	return this
}

// GherkinParser tokens.
const (
	GherkinParserEOF           = antlr.TokenEOF
	GherkinParserSTEP_KW       = 1
	GherkinParserFEATURE_KW    = 2
	GherkinParserBACKGROUND_KW = 3
	GherkinParserSCENARIO_KW   = 4
	GherkinParserOUTLINE_KW    = 5
	GherkinParserEXAMPLES_KW   = 6
	GherkinParserEOL           = 7
	GherkinParserWS            = 8
	GherkinParserPIPE          = 9
	GherkinParserESCAPED_PIPE  = 10
	GherkinParserCOLON         = 11
	GherkinParserAT            = 12
	GherkinParserHASH          = 13
	GherkinParserTRIPLE_QUOTE  = 14
	GherkinParserQUOTE         = 15
	GherkinParserCHAR          = 16
)

// GherkinParser rules.
const (
	GherkinParserRULE_tagName          = 0
	GherkinParserRULE_tag              = 1
	GherkinParserRULE_lineContent      = 2
	GherkinParserRULE_comment          = 3
	GherkinParserRULE_annotation       = 4
	GherkinParserRULE_trailingComment  = 5
	GherkinParserRULE_descriptionLine  = 6
	GherkinParserRULE_name             = 7
	GherkinParserRULE_feature          = 8
	GherkinParserRULE_background       = 9
	GherkinParserRULE_abstractScenario = 10
	GherkinParserRULE_scenario         = 11
	GherkinParserRULE_outline          = 12
	GherkinParserRULE_examples         = 13
	GherkinParserRULE_step             = 14
	GherkinParserRULE_row              = 15
	GherkinParserRULE_cell             = 16
	GherkinParserRULE_document         = 17
	GherkinParserRULE_documentIndent   = 18
	GherkinParserRULE_documentContent  = 19
)

// ITagNameContext is an interface to support dynamic dispatch.
type ITagNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagNameContext differentiates from other interfaces.
	IsTagNameContext()
}

type TagNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagNameContext() *TagNameContext {
	var p = new(TagNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_tagName
	return p
}

func (*TagNameContext) IsTagNameContext() {}

func NewTagNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagNameContext {
	var p = new(TagNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_tagName

	return p
}

func (s *TagNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TagNameContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *TagNameContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *TagNameContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *TagNameContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *TagNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterTagName(s)
	}
}

func (s *TagNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitTagName(s)
	}
}

func (p *GherkinParser) TagName() (localctx ITagNameContext) {
	localctx = NewTagNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GherkinParserRULE_tagName)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GherkinParserSTEP_KW)|(1<<GherkinParserFEATURE_KW)|(1<<GherkinParserBACKGROUND_KW)|(1<<GherkinParserSCENARIO_KW)|(1<<GherkinParserOUTLINE_KW)|(1<<GherkinParserEXAMPLES_KW)|(1<<GherkinParserPIPE)|(1<<GherkinParserESCAPED_PIPE)|(1<<GherkinParserCOLON)|(1<<GherkinParserAT)|(1<<GherkinParserHASH)|(1<<GherkinParserTRIPLE_QUOTE)|(1<<GherkinParserQUOTE)|(1<<GherkinParserCHAR))) != 0) {
		p.SetState(40)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == GherkinParserEOL || _la == GherkinParserWS {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) AT() antlr.TerminalNode {
	return s.GetToken(GherkinParserAT, 0)
}

func (s *TagContext) TagName() ITagNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagNameContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *GherkinParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GherkinParserRULE_tag)

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
		p.SetState(45)
		p.Match(GherkinParserAT)
	}
	{
		p.SetState(46)
		p.TagName()
	}

	return localctx
}

// ILineContentContext is an interface to support dynamic dispatch.
type ILineContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContentContext differentiates from other interfaces.
	IsLineContentContext()
}

type LineContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContentContext() *LineContentContext {
	var p = new(LineContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_lineContent
	return p
}

func (*LineContentContext) IsLineContentContext() {}

func NewLineContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContentContext {
	var p = new(LineContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_lineContent

	return p
}

func (s *LineContentContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContentContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *LineContentContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *LineContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterLineContent(s)
	}
}

func (s *LineContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitLineContent(s)
	}
}

func (p *GherkinParser) LineContent() (localctx ILineContentContext) {
	localctx = NewLineContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GherkinParserRULE_lineContent)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GherkinParserSTEP_KW)|(1<<GherkinParserFEATURE_KW)|(1<<GherkinParserBACKGROUND_KW)|(1<<GherkinParserSCENARIO_KW)|(1<<GherkinParserOUTLINE_KW)|(1<<GherkinParserEXAMPLES_KW)|(1<<GherkinParserWS)|(1<<GherkinParserPIPE)|(1<<GherkinParserESCAPED_PIPE)|(1<<GherkinParserCOLON)|(1<<GherkinParserAT)|(1<<GherkinParserHASH)|(1<<GherkinParserTRIPLE_QUOTE)|(1<<GherkinParserQUOTE)|(1<<GherkinParserCHAR))) != 0) {
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == GherkinParserEOL {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(GherkinParserHASH, 0)
}

func (s *CommentContext) LineContent() ILineContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContentContext)
}

func (s *CommentContext) EOL() antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, 0)
}

func (s *CommentContext) WS() antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *GherkinParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GherkinParserRULE_comment)
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
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(53)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(56)
		p.Match(GherkinParserHASH)
	}
	{
		p.SetState(57)
		p.LineContent()
	}
	{
		p.SetState(58)
		p.Match(GherkinParserEOL)
	}

	return localctx
}

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *AnnotationContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *AnnotationContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *AnnotationContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *AnnotationContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *AnnotationContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *AnnotationContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *GherkinParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GherkinParserRULE_annotation)
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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Comment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GherkinParserWS {
			{
				p.SetState(61)
				p.Match(GherkinParserWS)
			}

		}
		{
			p.SetState(64)
			p.Tag()
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(71)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case GherkinParserWS:
					{
						p.SetState(65)
						p.Match(GherkinParserWS)
					}

				case GherkinParserEOL:
					p.SetState(67)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for ok := true; ok; ok = _la == GherkinParserEOL {
						{
							p.SetState(66)
							p.Match(GherkinParserEOL)
						}

						p.SetState(69)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}
				{
					p.SetState(73)
					p.Tag()
				}

			}
			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GherkinParserWS {
			{
				p.SetState(79)
				p.Match(GherkinParserWS)
			}

			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(85)
			p.Match(GherkinParserEOL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GherkinParserWS {
			{
				p.SetState(87)
				p.Match(GherkinParserWS)
			}

		}
		{
			p.SetState(90)
			p.Match(GherkinParserEOL)
		}

	}

	return localctx
}

// ITrailingCommentContext is an interface to support dynamic dispatch.
type ITrailingCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrailingCommentContext differentiates from other interfaces.
	IsTrailingCommentContext()
}

type TrailingCommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrailingCommentContext() *TrailingCommentContext {
	var p = new(TrailingCommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_trailingComment
	return p
}

func (*TrailingCommentContext) IsTrailingCommentContext() {}

func NewTrailingCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrailingCommentContext {
	var p = new(TrailingCommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_trailingComment

	return p
}

func (s *TrailingCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *TrailingCommentContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *TrailingCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrailingCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrailingCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterTrailingComment(s)
	}
}

func (s *TrailingCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitTrailingComment(s)
	}
}

func (p *GherkinParser) TrailingComment() (localctx ITrailingCommentContext) {
	localctx = NewTrailingCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GherkinParserRULE_trailingComment)

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
		p.SetState(93)
		p.Comment()
	}

	return localctx
}

// IDescriptionLineContext is an interface to support dynamic dispatch.
type IDescriptionLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionLineContext differentiates from other interfaces.
	IsDescriptionLineContext()
}

type DescriptionLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionLineContext() *DescriptionLineContext {
	var p = new(DescriptionLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_descriptionLine
	return p
}

func (*DescriptionLineContext) IsDescriptionLineContext() {}

func NewDescriptionLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionLineContext {
	var p = new(DescriptionLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_descriptionLine

	return p
}

func (s *DescriptionLineContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionLineContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *DescriptionLineContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *DescriptionLineContext) FEATURE_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserFEATURE_KW, 0)
}

func (s *DescriptionLineContext) BACKGROUND_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserBACKGROUND_KW, 0)
}

func (s *DescriptionLineContext) SCENARIO_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserSCENARIO_KW, 0)
}

func (s *DescriptionLineContext) OUTLINE_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserOUTLINE_KW, 0)
}

func (s *DescriptionLineContext) STEP_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserSTEP_KW, 0)
}

func (s *DescriptionLineContext) EXAMPLES_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserEXAMPLES_KW, 0)
}

func (s *DescriptionLineContext) AT() antlr.TerminalNode {
	return s.GetToken(GherkinParserAT, 0)
}

func (s *DescriptionLineContext) HASH() antlr.TerminalNode {
	return s.GetToken(GherkinParserHASH, 0)
}

func (s *DescriptionLineContext) PIPE() antlr.TerminalNode {
	return s.GetToken(GherkinParserPIPE, 0)
}

func (s *DescriptionLineContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *DescriptionLineContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *DescriptionLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterDescriptionLine(s)
	}
}

func (s *DescriptionLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitDescriptionLine(s)
	}
}

func (p *GherkinParser) DescriptionLine() (localctx IDescriptionLineContext) {
	localctx = NewDescriptionLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GherkinParserRULE_descriptionLine)
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

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GherkinParserWS {
			{
				p.SetState(95)
				p.Match(GherkinParserWS)
			}

		}
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GherkinParserSTEP_KW)|(1<<GherkinParserFEATURE_KW)|(1<<GherkinParserBACKGROUND_KW)|(1<<GherkinParserSCENARIO_KW)|(1<<GherkinParserOUTLINE_KW)|(1<<GherkinParserEXAMPLES_KW)|(1<<GherkinParserEOL)|(1<<GherkinParserWS)|(1<<GherkinParserPIPE)|(1<<GherkinParserAT)|(1<<GherkinParserHASH))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GherkinParserSTEP_KW)|(1<<GherkinParserFEATURE_KW)|(1<<GherkinParserBACKGROUND_KW)|(1<<GherkinParserSCENARIO_KW)|(1<<GherkinParserOUTLINE_KW)|(1<<GherkinParserEXAMPLES_KW)|(1<<GherkinParserWS)|(1<<GherkinParserPIPE)|(1<<GherkinParserESCAPED_PIPE)|(1<<GherkinParserCOLON)|(1<<GherkinParserAT)|(1<<GherkinParserHASH)|(1<<GherkinParserTRIPLE_QUOTE)|(1<<GherkinParserQUOTE)|(1<<GherkinParserCHAR))) != 0 {
			p.SetState(99)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == GherkinParserEOL {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(105)
			p.Match(GherkinParserEOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GherkinParserWS {
			{
				p.SetState(106)
				p.Match(GherkinParserWS)
			}

		}
		{
			p.SetState(109)
			p.Match(GherkinParserEOL)
		}

	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) AllDescriptionLine() []IDescriptionLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDescriptionLineContext)(nil)).Elem())
	var tst = make([]IDescriptionLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDescriptionLineContext)
		}
	}

	return tst
}

func (s *NameContext) DescriptionLine(i int) IDescriptionLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDescriptionLineContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *GherkinParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GherkinParserRULE_name)

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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(112)
				p.DescriptionLine()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IFeatureContext is an interface to support dynamic dispatch.
type IFeatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFeatureContext differentiates from other interfaces.
	IsFeatureContext()
}

type FeatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeatureContext() *FeatureContext {
	var p = new(FeatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_feature
	return p
}

func (*FeatureContext) IsFeatureContext() {}

func NewFeatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeatureContext {
	var p = new(FeatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_feature

	return p
}

func (s *FeatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FeatureContext) FEATURE_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserFEATURE_KW, 0)
}

func (s *FeatureContext) COLON() antlr.TerminalNode {
	return s.GetToken(GherkinParserCOLON, 0)
}

func (s *FeatureContext) LineContent() ILineContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContentContext)
}

func (s *FeatureContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *FeatureContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *FeatureContext) EOF() antlr.TerminalNode {
	return s.GetToken(GherkinParserEOF, 0)
}

func (s *FeatureContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *FeatureContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FeatureContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *FeatureContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *FeatureContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FeatureContext) Background() IBackgroundContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBackgroundContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBackgroundContext)
}

func (s *FeatureContext) AllAbstractScenario() []IAbstractScenarioContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAbstractScenarioContext)(nil)).Elem())
	var tst = make([]IAbstractScenarioContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAbstractScenarioContext)
		}
	}

	return tst
}

func (s *FeatureContext) AbstractScenario(i int) IAbstractScenarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbstractScenarioContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAbstractScenarioContext)
}

func (s *FeatureContext) AllTrailingComment() []ITrailingCommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITrailingCommentContext)(nil)).Elem())
	var tst = make([]ITrailingCommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITrailingCommentContext)
		}
	}

	return tst
}

func (s *FeatureContext) TrailingComment(i int) ITrailingCommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrailingCommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITrailingCommentContext)
}

func (s *FeatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterFeature(s)
	}
}

func (s *FeatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitFeature(s)
	}
}

func (p *GherkinParser) Feature() (localctx IFeatureContext) {
	localctx = NewFeatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GherkinParserRULE_feature)
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
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(117)
				p.Annotation()
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(123)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(126)
		p.Match(GherkinParserFEATURE_KW)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(127)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(130)
		p.Match(GherkinParserCOLON)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(131)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(134)
		p.LineContent()
	}
	{
		p.SetState(135)
		p.Match(GherkinParserEOL)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(136)
			p.Name()
		}

	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(139)
			p.Background()
		}

	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(142)
				p.AbstractScenario()
			}

		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GherkinParserEOL)|(1<<GherkinParserWS)|(1<<GherkinParserHASH))) != 0 {
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == GherkinParserWS {
				{
					p.SetState(148)
					p.Match(GherkinParserWS)
				}

			}
			{
				p.SetState(151)
				p.Match(GherkinParserEOL)
			}

		case 2:
			{
				p.SetState(152)
				p.TrailingComment()
			}

		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(158)
		p.Match(GherkinParserEOF)
	}

	return localctx
}

// IBackgroundContext is an interface to support dynamic dispatch.
type IBackgroundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBackgroundContext differentiates from other interfaces.
	IsBackgroundContext()
}

type BackgroundContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBackgroundContext() *BackgroundContext {
	var p = new(BackgroundContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_background
	return p
}

func (*BackgroundContext) IsBackgroundContext() {}

func NewBackgroundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BackgroundContext {
	var p = new(BackgroundContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_background

	return p
}

func (s *BackgroundContext) GetParser() antlr.Parser { return s.parser }

func (s *BackgroundContext) BACKGROUND_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserBACKGROUND_KW, 0)
}

func (s *BackgroundContext) COLON() antlr.TerminalNode {
	return s.GetToken(GherkinParserCOLON, 0)
}

func (s *BackgroundContext) LineContent() ILineContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContentContext)
}

func (s *BackgroundContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *BackgroundContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *BackgroundContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *BackgroundContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *BackgroundContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *BackgroundContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *BackgroundContext) AllStep() []IStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStepContext)(nil)).Elem())
	var tst = make([]IStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStepContext)
		}
	}

	return tst
}

func (s *BackgroundContext) Step(i int) IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *BackgroundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BackgroundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BackgroundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterBackground(s)
	}
}

func (s *BackgroundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitBackground(s)
	}
}

func (p *GherkinParser) Background() (localctx IBackgroundContext) {
	localctx = NewBackgroundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GherkinParserRULE_background)
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
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(160)
				p.Annotation()
			}

		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(166)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(169)
		p.Match(GherkinParserBACKGROUND_KW)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(170)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(173)
		p.Match(GherkinParserCOLON)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(174)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(177)
		p.LineContent()
	}
	{
		p.SetState(178)
		p.Match(GherkinParserEOL)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(179)
					p.Step()
				}

			case 2:
				p.SetState(181)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == GherkinParserWS {
					{
						p.SetState(180)
						p.Match(GherkinParserWS)
					}

				}
				{
					p.SetState(183)
					p.Match(GherkinParserEOL)
				}

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// IAbstractScenarioContext is an interface to support dynamic dispatch.
type IAbstractScenarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbstractScenarioContext differentiates from other interfaces.
	IsAbstractScenarioContext()
}

type AbstractScenarioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbstractScenarioContext() *AbstractScenarioContext {
	var p = new(AbstractScenarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_abstractScenario
	return p
}

func (*AbstractScenarioContext) IsAbstractScenarioContext() {}

func NewAbstractScenarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbstractScenarioContext {
	var p = new(AbstractScenarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_abstractScenario

	return p
}

func (s *AbstractScenarioContext) GetParser() antlr.Parser { return s.parser }

func (s *AbstractScenarioContext) Scenario() IScenarioContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScenarioContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScenarioContext)
}

func (s *AbstractScenarioContext) Outline() IOutlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutlineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutlineContext)
}

func (s *AbstractScenarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbstractScenarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbstractScenarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterAbstractScenario(s)
	}
}

func (s *AbstractScenarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitAbstractScenario(s)
	}
}

func (p *GherkinParser) AbstractScenario() (localctx IAbstractScenarioContext) {
	localctx = NewAbstractScenarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GherkinParserRULE_abstractScenario)

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

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Scenario()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Outline()
		}

	}

	return localctx
}

// IScenarioContext is an interface to support dynamic dispatch.
type IScenarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScenarioContext differentiates from other interfaces.
	IsScenarioContext()
}

type ScenarioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScenarioContext() *ScenarioContext {
	var p = new(ScenarioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_scenario
	return p
}

func (*ScenarioContext) IsScenarioContext() {}

func NewScenarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScenarioContext {
	var p = new(ScenarioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_scenario

	return p
}

func (s *ScenarioContext) GetParser() antlr.Parser { return s.parser }

func (s *ScenarioContext) SCENARIO_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserSCENARIO_KW, 0)
}

func (s *ScenarioContext) COLON() antlr.TerminalNode {
	return s.GetToken(GherkinParserCOLON, 0)
}

func (s *ScenarioContext) LineContent() ILineContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContentContext)
}

func (s *ScenarioContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *ScenarioContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *ScenarioContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ScenarioContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ScenarioContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *ScenarioContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *ScenarioContext) AllStep() []IStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStepContext)(nil)).Elem())
	var tst = make([]IStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStepContext)
		}
	}

	return tst
}

func (s *ScenarioContext) Step(i int) IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *ScenarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScenarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScenarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterScenario(s)
	}
}

func (s *ScenarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitScenario(s)
	}
}

func (p *GherkinParser) Scenario() (localctx IScenarioContext) {
	localctx = NewScenarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GherkinParserRULE_scenario)
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
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(192)
				p.Annotation()
			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(198)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(201)
		p.Match(GherkinParserSCENARIO_KW)
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(202)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(205)
		p.Match(GherkinParserCOLON)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(206)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(209)
		p.LineContent()
	}
	{
		p.SetState(210)
		p.Match(GherkinParserEOL)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(211)
					p.Step()
				}

			case 2:
				p.SetState(213)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == GherkinParserWS {
					{
						p.SetState(212)
						p.Match(GherkinParserWS)
					}

				}
				{
					p.SetState(215)
					p.Match(GherkinParserEOL)
				}

			}

		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IOutlineContext is an interface to support dynamic dispatch.
type IOutlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutlineContext differentiates from other interfaces.
	IsOutlineContext()
}

type OutlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutlineContext() *OutlineContext {
	var p = new(OutlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_outline
	return p
}

func (*OutlineContext) IsOutlineContext() {}

func NewOutlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutlineContext {
	var p = new(OutlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_outline

	return p
}

func (s *OutlineContext) GetParser() antlr.Parser { return s.parser }

func (s *OutlineContext) OUTLINE_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserOUTLINE_KW, 0)
}

func (s *OutlineContext) COLON() antlr.TerminalNode {
	return s.GetToken(GherkinParserCOLON, 0)
}

func (s *OutlineContext) LineContent() ILineContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContentContext)
}

func (s *OutlineContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *OutlineContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *OutlineContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *OutlineContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *OutlineContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *OutlineContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *OutlineContext) AllStep() []IStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStepContext)(nil)).Elem())
	var tst = make([]IStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStepContext)
		}
	}

	return tst
}

func (s *OutlineContext) Step(i int) IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *OutlineContext) AllExamples() []IExamplesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExamplesContext)(nil)).Elem())
	var tst = make([]IExamplesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExamplesContext)
		}
	}

	return tst
}

func (s *OutlineContext) Examples(i int) IExamplesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExamplesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExamplesContext)
}

func (s *OutlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterOutline(s)
	}
}

func (s *OutlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitOutline(s)
	}
}

func (p *GherkinParser) Outline() (localctx IOutlineContext) {
	localctx = NewOutlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GherkinParserRULE_outline)
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
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(221)
				p.Annotation()
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(227)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(230)
		p.Match(GherkinParserOUTLINE_KW)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(231)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(234)
		p.Match(GherkinParserCOLON)
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(235)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(238)
		p.LineContent()
	}
	{
		p.SetState(239)
		p.Match(GherkinParserEOL)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(240)
					p.Step()
				}

			case 2:
				p.SetState(242)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == GherkinParserWS {
					{
						p.SetState(241)
						p.Match(GherkinParserWS)
					}

				}
				{
					p.SetState(244)
					p.Match(GherkinParserEOL)
				}

			}

		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(250)
				p.Examples()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == GherkinParserWS {
				{
					p.SetState(255)
					p.Match(GherkinParserWS)
				}

			}
			{
				p.SetState(258)
				p.Match(GherkinParserEOL)
			}

		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext())
	}

	return localctx
}

// IExamplesContext is an interface to support dynamic dispatch.
type IExamplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExamplesContext differentiates from other interfaces.
	IsExamplesContext()
}

type ExamplesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExamplesContext() *ExamplesContext {
	var p = new(ExamplesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_examples
	return p
}

func (*ExamplesContext) IsExamplesContext() {}

func NewExamplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExamplesContext {
	var p = new(ExamplesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_examples

	return p
}

func (s *ExamplesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExamplesContext) EXAMPLES_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserEXAMPLES_KW, 0)
}

func (s *ExamplesContext) COLON() antlr.TerminalNode {
	return s.GetToken(GherkinParserCOLON, 0)
}

func (s *ExamplesContext) EOL() antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, 0)
}

func (s *ExamplesContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ExamplesContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ExamplesContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *ExamplesContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *ExamplesContext) LineContent() ILineContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContentContext)
}

func (s *ExamplesContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *ExamplesContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *ExamplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExamplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExamplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterExamples(s)
	}
}

func (s *ExamplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitExamples(s)
	}
}

func (p *GherkinParser) Examples() (localctx IExamplesContext) {
	localctx = NewExamplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GherkinParserRULE_examples)
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
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(264)
				p.Annotation()
			}

		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(270)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(273)
		p.Match(GherkinParserEXAMPLES_KW)
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(274)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(277)
		p.Match(GherkinParserCOLON)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(278)
			p.Match(GherkinParserWS)
		}

	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GherkinParserSTEP_KW)|(1<<GherkinParserFEATURE_KW)|(1<<GherkinParserBACKGROUND_KW)|(1<<GherkinParserSCENARIO_KW)|(1<<GherkinParserOUTLINE_KW)|(1<<GherkinParserEXAMPLES_KW)|(1<<GherkinParserWS)|(1<<GherkinParserPIPE)|(1<<GherkinParserESCAPED_PIPE)|(1<<GherkinParserCOLON)|(1<<GherkinParserAT)|(1<<GherkinParserHASH)|(1<<GherkinParserTRIPLE_QUOTE)|(1<<GherkinParserQUOTE)|(1<<GherkinParserCHAR))) != 0 {
		{
			p.SetState(281)
			p.LineContent()
		}

	}
	{
		p.SetState(284)
		p.Match(GherkinParserEOL)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(285)
				p.Row()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())
	}

	return localctx
}

// IStepContext is an interface to support dynamic dispatch.
type IStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepContext differentiates from other interfaces.
	IsStepContext()
}

type StepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepContext() *StepContext {
	var p = new(StepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_step
	return p
}

func (*StepContext) IsStepContext() {}

func NewStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepContext {
	var p = new(StepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_step

	return p
}

func (s *StepContext) GetParser() antlr.Parser { return s.parser }

func (s *StepContext) STEP_KW() antlr.TerminalNode {
	return s.GetToken(GherkinParserSTEP_KW, 0)
}

func (s *StepContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *StepContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *StepContext) LineContent() ILineContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContentContext)
}

func (s *StepContext) EOL() antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, 0)
}

func (s *StepContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *StepContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *StepContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *StepContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *StepContext) AllDocument() []IDocumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDocumentContext)(nil)).Elem())
	var tst = make([]IDocumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDocumentContext)
		}
	}

	return tst
}

func (s *StepContext) Document(i int) IDocumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *StepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterStep(s)
	}
}

func (s *StepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitStep(s)
	}
}

func (p *GherkinParser) Step() (localctx IStepContext) {
	localctx = NewStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GherkinParserRULE_step)
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
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(290)
				p.Annotation()
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(296)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(299)
		p.Match(GherkinParserSTEP_KW)
	}
	{
		p.SetState(300)
		p.Match(GherkinParserWS)
	}
	{
		p.SetState(301)
		p.LineContent()
	}
	{
		p.SetState(302)
		p.Match(GherkinParserEOL)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(303)
					p.Row()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(306)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext())
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 2 {
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(308)
					p.Document()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserPIPE)
}

func (s *RowContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserPIPE, i)
}

func (s *RowContext) EOL() antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, 0)
}

func (s *RowContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *RowContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *RowContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserWS)
}

func (s *RowContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, i)
}

func (s *RowContext) AllCell() []ICellContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICellContext)(nil)).Elem())
	var tst = make([]ICellContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICellContext)
		}
	}

	return tst
}

func (s *RowContext) Cell(i int) ICellContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICellContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICellContext)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitRow(s)
	}
}

func (p *GherkinParser) Row() (localctx IRowContext) {
	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GherkinParserRULE_row)
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
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(315)
				p.Annotation()
			}

		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext())
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(321)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(324)
		p.Match(GherkinParserPIPE)
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(325)
				p.Cell()
			}
			{
				p.SetState(326)
				p.Match(GherkinParserPIPE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(332)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(335)
		p.Match(GherkinParserEOL)
	}

	return localctx
}

// ICellContext is an interface to support dynamic dispatch.
type ICellContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCellContext differentiates from other interfaces.
	IsCellContext()
}

type CellContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCellContext() *CellContext {
	var p = new(CellContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_cell
	return p
}

func (*CellContext) IsCellContext() {}

func NewCellContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CellContext {
	var p = new(CellContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_cell

	return p
}

func (s *CellContext) GetParser() antlr.Parser { return s.parser }

func (s *CellContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserPIPE)
}

func (s *CellContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserPIPE, i)
}

func (s *CellContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserEOL)
}

func (s *CellContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, i)
}

func (s *CellContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CellContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CellContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterCell(s)
	}
}

func (s *CellContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitCell(s)
	}
}

func (p *GherkinParser) Cell() (localctx ICellContext) {
	localctx = NewCellContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GherkinParserRULE_cell)
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GherkinParserSTEP_KW)|(1<<GherkinParserFEATURE_KW)|(1<<GherkinParserBACKGROUND_KW)|(1<<GherkinParserSCENARIO_KW)|(1<<GherkinParserOUTLINE_KW)|(1<<GherkinParserEXAMPLES_KW)|(1<<GherkinParserWS)|(1<<GherkinParserESCAPED_PIPE)|(1<<GherkinParserCOLON)|(1<<GherkinParserAT)|(1<<GherkinParserHASH)|(1<<GherkinParserTRIPLE_QUOTE)|(1<<GherkinParserQUOTE)|(1<<GherkinParserCHAR))) != 0 {
		p.SetState(337)
		_la = p.GetTokenStream().LA(1)

		if _la <= 0 || _la == GherkinParserEOL || _la == GherkinParserPIPE {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) DocumentIndent() IDocumentIndentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentIndentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentIndentContext)
}

func (s *DocumentContext) AllTRIPLE_QUOTE() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserTRIPLE_QUOTE)
}

func (s *DocumentContext) TRIPLE_QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserTRIPLE_QUOTE, i)
}

func (s *DocumentContext) DocumentContent() IDocumentContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentContentContext)
}

func (s *DocumentContext) EOL() antlr.TerminalNode {
	return s.GetToken(GherkinParserEOL, 0)
}

func (s *DocumentContext) WS() antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, 0)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *GherkinParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GherkinParserRULE_document)
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
		p.SetState(343)
		p.DocumentIndent()
	}
	{
		p.SetState(344)
		p.Match(GherkinParserTRIPLE_QUOTE)
	}
	{
		p.SetState(345)
		p.DocumentContent()
	}
	{
		p.SetState(346)
		p.Match(GherkinParserTRIPLE_QUOTE)
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(347)
			p.Match(GherkinParserWS)
		}

	}
	{
		p.SetState(350)
		p.Match(GherkinParserEOL)
	}

	return localctx
}

// IDocumentIndentContext is an interface to support dynamic dispatch.
type IDocumentIndentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentIndentContext differentiates from other interfaces.
	IsDocumentIndentContext()
}

type DocumentIndentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentIndentContext() *DocumentIndentContext {
	var p = new(DocumentIndentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_documentIndent
	return p
}

func (*DocumentIndentContext) IsDocumentIndentContext() {}

func NewDocumentIndentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentIndentContext {
	var p = new(DocumentIndentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_documentIndent

	return p
}

func (s *DocumentIndentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentIndentContext) WS() antlr.TerminalNode {
	return s.GetToken(GherkinParserWS, 0)
}

func (s *DocumentIndentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentIndentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentIndentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterDocumentIndent(s)
	}
}

func (s *DocumentIndentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitDocumentIndent(s)
	}
}

func (p *GherkinParser) DocumentIndent() (localctx IDocumentIndentContext) {
	localctx = NewDocumentIndentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GherkinParserRULE_documentIndent)
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
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GherkinParserWS {
		{
			p.SetState(352)
			p.Match(GherkinParserWS)
		}

	}

	return localctx
}

// IDocumentContentContext is an interface to support dynamic dispatch.
type IDocumentContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContentContext differentiates from other interfaces.
	IsDocumentContentContext()
}

type DocumentContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContentContext() *DocumentContentContext {
	var p = new(DocumentContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GherkinParserRULE_documentContent
	return p
}

func (*DocumentContentContext) IsDocumentContentContext() {}

func NewDocumentContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContentContext {
	var p = new(DocumentContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GherkinParserRULE_documentContent

	return p
}

func (s *DocumentContentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContentContext) AllTRIPLE_QUOTE() []antlr.TerminalNode {
	return s.GetTokens(GherkinParserTRIPLE_QUOTE)
}

func (s *DocumentContentContext) TRIPLE_QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(GherkinParserTRIPLE_QUOTE, i)
}

func (s *DocumentContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.EnterDocumentContent(s)
	}
}

func (s *DocumentContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GherkinListener); ok {
		listenerT.ExitDocumentContent(s)
	}
}

func (p *GherkinParser) DocumentContent() (localctx IDocumentContentContext) {
	localctx = NewDocumentContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GherkinParserRULE_documentContent)
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
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(355)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == GherkinParserTRIPLE_QUOTE {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())
	}

	return localctx
}
