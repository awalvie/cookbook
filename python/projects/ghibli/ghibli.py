#!/bin/python3

import requests
import os
import shutil
from bs4 import BeautifulSoup

def filmList(url):
        film_list = []

        page = requests.get(url)
        soup = BeautifulSoup(page.text, "html.parser")

        links = soup.find_all('a', class_='panelarea')

        for link in links:
                film_url = link.get('href')
                film_title = film_url.split('/')[4]
                
                film_list.append([film_title, film_url])


        return film_list

def imageList(title, url):
        image_list = []

        page = requests.get(url)
        soup = BeautifulSoup(page.text, "html.parser")

        links = soup.find_all('a', class_="panelarea")

        for link in links:
                image_url = link.get('href')
                image_list.append(image_url)

        return [title, image_list]

def downloadImages(image_list):
        for title in image_list:
                os.makedirs("./images/" + title[0])
                os.chdir("./images/" + title[0])

                for image_url in title[1]:
                    r = requests.get(image_url, stream=True)
                    filename = image_url.split("/")[-1]

                    print("Downloading: ", filename)
                    with open(filename,'wb') as f:
                            shutil.copyfileobj(r.raw, f)
                
                os.chdir("../../")


if __name__ == "__main__":
    image_list = []
    
    url = "https://www.ghibli.jp/info/013409/"
    film_list = filmList(url)
    
    for film in film_list:
            lst = imageList(film[0], film[1])
            image_list.append(lst)

    downloadImages(image_list)
