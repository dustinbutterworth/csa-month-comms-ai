# ChatGPT Cybersecurity Awareness Month Communications Creator

A GoLang project that reads a file, generates a blog post with ChatGPT, and writes the results to a file.

This is designed to generate content for Cybersecurity Awareness Month for teams 
that need to get some communications out but may be short staffed and don't have 
the manpower or writing experience to do it manually. 

## Overview

This project is designed to help you automate the process of generating blog posts on a given topic by using OpenAI's ChatGPT. It consists of a GoLang script that reads the contents of a file, sends a request to ChatGPT to generate a blog post based on the file's content, and then writes the generated blog post to a file.

## Prompt

The prompt I'm using to generate the report is based off the [Cybersecuirty Awareness Month Partner Toolkit](https://www.cisa.gov/sites/default/files/2023-09/Cybersecurity%20Awareness%20Month%202023%20Toolkit%20Guide%20FINAL_508c.pdf) presented by CISA for October 2023:

```
You are an Cyber Threat Intelligence Analyst.
This month is Cybersecurity Awareness month. This month you will be writing blog posts
for the entire company on any topic I provide for you. 
I will give you a topic, and you will write a blog post 800 words or less in length
that is targeted toward a general audience in a technology company.
You you should use the following tone of voice:
1. Positivity – Scare tactics don’t work. Instead of using scary “hackers in hoodies” imagery, talk
about the benefits of reducing cyber risks and how strengthening cybersecurity can protect
what matters most in our lives.
2. Approachability – Cybersecurity seems like a complex subject to many, but it’s really all about
people. Make cybersecurity relatable and share how practicing good cyber hygiene is
something that anyone can do.
3. Simplicity – Avoid jargon and be sure to define acronyms
4. Back to basics – Even just practicing cyber hygiene basics provides a solid baseline of security. 
```

I recommend you copy and paste the individual "Key Behaviors" content from the 
partner toolkit into their own separate files and use those for your inputs. 
However, any topic can be used.

## Prerequisites

Before using this project, make sure you have the following prerequisites:

- GoLang installed on your system.
- An OpenAI API key. You can obtain one by signing up for the OpenAI API.

## Usage

1. Clone the repository to your local machine:

   ```shell
   git clone <repository_url>
   cd <repository_name>
   ```
2. Set your OpenAI API key as an environment variable:

    ```shell 
    export OPENAI_API_KEY=your_api_key_here
    ```

3. Run the program with the path to the file you want to use as input:

    ```shell 
    go run main.go path/to/your/file.txt
    ```

The program will generate a blog post based on the contents of the file and write it to the "articles" directory.

## Configuration 

The program will generate a blog post based on the contents of the file and write it to the "articles" directory.

## Acknowledgements

* OpenAI for providing the ChatGPT API.

## Author 

* Dustin Butterworth
