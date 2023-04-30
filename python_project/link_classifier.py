# link_classifier.py
import requests
from bs4 import BeautifulSoup
from urllib.parse import urlparse, urljoin


class LinkClassifier:
    def __init__(self, url):
        self.url = url

    def classify_links(self):
        response = requests.get(self.url)
        soup = BeautifulSoup(response.text, 'html.parser')

        domain = urlparse(self.url).netloc
        links = soup.find_all('a', href=True)

        inbound_links = []
        outbound_links = []
        for link in links:
            href = link['href']
            target_domain = urlparse(href).netloc

            # 絶対URLに変換
            absolute_url = urljoin(self.url, href)

            if not target_domain:
                inbound_links.append(absolute_url)
            elif domain == target_domain:
                inbound_links.append(absolute_url)
            else:
                outbound_links.append(absolute_url)

        return {
            'inbound_links': inbound_links,
            'outbound_links': outbound_links
        }

    def check_reprint(self, original_website_url):
        response = requests.get(self.url)
        soup = BeautifulSoup(response.text, 'html.parser')

        original_response = requests.get(original_website_url)
        original_soup = BeautifulSoup(original_response.text, 'html.parser')

        reprint = soup.text == original_soup.text

        return reprint
