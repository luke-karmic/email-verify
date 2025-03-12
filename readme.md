<div align="center">

  [![Unlicense License][license-shield]][license-url]
  [![LinkedIn][linkedin-shield]][linkedin-url]

</div>


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">Email verifier - Lead verification</h3>

  <p align="center">
    A Docker-based quick GoLang lead verifier, instead of paying monthly for email verification, use TrueMail Lib to do it manually.
    Shout out to @bestwebua and @RichiMaulana for the TrueMail Lib allowing this to work so easily.
    This requires SMTP port 25 to be open, so deploy this on a VPS in order for it to work.
    Just drop your CSV of leads into a file, give the path and run this on your server, it will output the success/failures & (%).
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

* Can add CSV file and run the program which will output the leads verified and the % of bounces
* Uses SMTP or MX to create an incomplete TCP request to verify the email

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

![Go][go-shield]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo
   ```sh
   https://github.com/luke-karmic/email-verify
   ```
2. Add your `leads.csv` file
3. `sudo apt-get install ccze` for the colorized logs
3. Update the docker file with your path to the CSV: `LEADS_FILE="./leads.csv"`
4. If you will use the `run.sh` script, you must give execution permissions by running `chmod +x run.sh`


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage with Docker (Manual)

1. run 
   ```sh
   build -t email-validator .
   ```
2. run 
   ```sh
   docker run -d --rm -v "./leads.csv:/app/leads.csv" email-validator
   ```
2. run 
   ```sh
   docker logs -f "$CONTAINER_ID" | ccze -A
   ```
3. View the logs output

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Create quick CSV Verifier


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Luke Taaffe

Project Link: [https://github.com/luke-karmic/email-verify](https://github.com/luke-karmic/email-verify)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/luketaaffe/
[go-shield]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge