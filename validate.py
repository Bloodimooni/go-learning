import requests
import random
#   Variables 

tokens = []



def randstr(lenn):
    alpha = "abcdefghijklmnopqrstuvwxyz0123456789"
    text = ''
    for i in range(0, lenn):
        text += alpha[random.randint(0, len(alpha) - 1)]
    return text


def check_tokens():
    with open("generated_tokens.txt","r+") as file:
        tokens = file.readlines()

    valid_token = 0
    invalid_token = 0
    valid_tokens = []
    for i in range(0,len(tokens)):
        changed = tokens[i].replace("\n","")
        changed2 = changed.replace("b'","")
        changed3 = changed2.replace("'","")
        header = {
                    "authority":            "discord.com",
                    "scheme":               "https",
                    "accept":               "*/*",
                    "accept-encoding":      "gzip, deflate, br",
                    "accept-language":      "en-US",
                    "Authorization":        f"{changed3}",
                    "content-length":       "0",
                    "cookie":               f"__cfuid={randstr(43)}; __dcfduid={randstr(32)}; locale=en-US",
                    "origin":               "https://discord.com",
                    "referer":              "https://discord.com/channels/@me",
                   "sec-fetch-dest":       "empty",
                    "sec-fetch-mode":       "cors",
                    "sec-fetch-site":       "same-origin",
                   "user-agent":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36",
                    "x-context-properties": "eyJsb2NhdGlvbiI6Ikludml0ZSBCdXR0b24gRW1iZWQiLCJsb2NhdGlvbl9ndWlsZF9pZCI6Ijg3OTc4MjM4MDAxMTk0NjAyNCIsImxvY2F0aW9uX2NoYW5uZWxfaWQiOiI4ODExMDg4MDc5NjE0MTk3OTYiLCJsb2NhdGlvbl9jaGFubmVsX3R5cGUiOjAsImxvY2F0aW9uX21lc3NhZ2VfaWQiOiI4ODExOTkzOTI5MTExNTkzNTcifQ==",
                    "x-debug-options":      "bugReporterEnabled",
                    "x-super-properties":   "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJjYW5hcnkiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC42MDAiLCJvc192ZXJzaW9uIjoiMTAuMC4yMjAwMCIsIm9zX2FyY2giOiJ4NjQiLCJzeXN0ZW1fbG9jYWxlIjoic2siLCJjbGllbnRfYnVpbGRfbnVtYmVyIjo5NTM1MywiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=",
                }

        r = requests.get("https://discordapp.com/api/v9/users/@me/library", headers=header)
        
        if r.status_code == 200:
            valid_token += 1
            valid_tokens.append(i)
        else:
            invalid_token += 1

    print("Valid tokens\t:\t{0}".format(valid_token))
    print("Invalid tokens\t:\t{0}".format(invalid_token))
    return valid_tokens


valid_tokens = check_tokens()
with open("valid_tokens.txt","w") as file:
    file.writelines(valid_tokens)
    