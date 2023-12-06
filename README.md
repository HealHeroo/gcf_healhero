# GCF HealHero

Repositori ini berisi proyek Google Cloud Functions untuk aplikasi HealHero.

## Instalasi

1. Clone repositori ini ke dalam direktori lokal Anda:

    ```bash
    git clone https://github.com/HealHero/gcf-healhero.git
    cd gcf-healhero
    ```

2. Instal dependensi proyek dengan perintah:

    ```bash
    go get -u ./...
    ```

## Struktur Direktori

Repositori ini terdiri dari 10 folder yang mewakili berbagai bagian dari aplikasi HealHero. Pada setiap file function.go memanggil dua baris impor paket, yaitu :
"github.com/GoogleCloudPlatform/functions-framework-go/functions"
- Paket ini terkait dengan Google Cloud Functions dan menyediakan fungsi-fungsi dan struktur data yang diperlukan untuk menulis fungsi Google Cloud Functions menggunakan Functions Framework for Go
"github.com/HealHeroo/be_healhero/module"
- paket ini akan berisi berbagai modul atau fungsi-fungsi yang diperlukan oleh proyek 


## Penggunaan

Untuk menjalankan fungsi-fungsi Google Cloud, ikuti petunjuk instalasi di atas. Setelah itu, Anda dapat mengimplementasikan dan menyinkronkan fungsi-fungsi tersebut dengan konsol Google Cloud.


