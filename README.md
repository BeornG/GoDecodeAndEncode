# **Hello**
A simple program using a command line interface to decode and encode

current issues:
parseint and parseuint cant handle longer numbers


--------------------ᶘ ᵒᴥᵒᶅ--------------------
dec
  -base2
        Decode base2(binary)
  -base64
        Decode base64
  -hex
        Decode hex
  -rot13
        Decode rot13, must be one string
enc
  -md5
        Encode md5, must be one string
  -rot13
        Encode rot13, must be one string
--------------------ᶘ ᵒᴥᵒᶅ--------------------


Example: go run . enc -rot13 "bongo bingo bengo"
Output: obatb ovatb oratb
