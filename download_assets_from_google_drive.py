import gdown
import os
import logging
logging.basicConfig(level=logging.DEBUG)

files = {
    'Charecter.nx' : 'https://drive.google.com/uc?id=18Ta2zVqfiPbENb_PKCGh2iOsM1v20rFq',
    'Sound.nx:' : 'https://drive.google.com/uc?id=1ZDpTkUbgPXT5au1p40oQB-NGUrXgCIFm',
    'String.nx' : 'https://drive.google.com/uc?id=1mbe_c6_WW2NyhgCRCFdHKmfzQmJoDPOz',
    'Item.nx' : 'https://drive.google.com/uc?id=13JCyBj5QNKvkzmsY-tBtEwM0025zTYen',
    'Map.nx' : 'https://drive.google.com/uc?id=1aur3HA2uVMBr934FZh63Ta689cyRyXwB',
    'Mob.nx' : 'https://drive.google.com/uc?id=1rHp455rzFGojCH8A1gnfOTurfkarOWce'
}

for name, url in files.items():
    gdown.download(url, os.path.join('assets', name), quiet=False)

# md5 = 'fa837a88f0c40c513d975104edf3da17'uc 
# gdown.cached_download(url, output, md5=md5, postprocess=gdown.extractall)