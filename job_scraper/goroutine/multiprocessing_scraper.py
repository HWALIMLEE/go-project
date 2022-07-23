import requests
from bs4 import BeautifulSoup
import multiprocessing
import csv
import time


class CrawlTest:
    def __init__(self):
        self.base_url = "https://kr.indeed.com/jobs?q=python&limit=50"

    def get_pages(self):
        res = requests.get(self.base_url)
        soup = BeautifulSoup(res.text, 'html.parser')
        pages = soup.find('div', {'class': 'pagination'})
        return len(pages.find_all('a'))


    def extract_job(self, card):
        title = card.find('h2', {'class': 'jobTitle'}).find('a').text
        name = card.find('span', {'class': 'companyName'}).text
        location = card.find('div', {'class': 'companyLocation'}).text
        return {
            'title': title,
            'name': name,
            'location': location
        }


    def get_page(self, page):
        page_url = self.base_url + "&start=" + str(page*50)
        res = requests.get(page_url)
        soup = BeautifulSoup(res.text, 'html.parser')
        search_cards = soup.find_all('td', {'class': 'resultContent'})
        result = [self.extract_job(card) for card in search_cards]
        return result

    def write_csv(self, result):
        keys = result[0].keys()

        with open('jobs.csv', 'w', newline='') as output_file:
            dict_writer = csv.DictWriter(output_file, keys)
            dict_writer.writeheader()
            dict_writer.writerows(result)


if __name__ == '__main__':
    start = time.time()
    crawl = CrawlTest()
    pages = list(range(1, crawl.get_pages()+1))
    pool = multiprocessing.Pool(processes=4)
    result = pool.map(crawl.get_page, [(page) for page in pages])
    final = sum(result, [])
    crawl.write_csv(final)
    pool.close()
    pool.join()
    end = time.time()
    print('time:', end-start)



