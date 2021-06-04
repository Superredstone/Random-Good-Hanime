from simple_image_download import simple_image_download as simp

downloads = int(input("Insert images to download: "))

response = simp.simple_image_download

# Do not change this
lst=["hentai"] 
for rep in lst:
    response().download(rep, downloads, ".jpeg")