# Dokumentasi Inventory Aplikasi

## FlowChart

```mermaid
flowchart TD
    A((Mulai)) --> B[Login]
    B -->|Autentikasi| C{Role}
    C -->|Admin| D[Lihat Kategori]
    C -->|Admin/Pengguna| E[Lihat Produk]
    C -->|Admin| F[Kelola Transaksi]


    E --> E1[Kelola Produk]
    E1 --> H[(Database: Produk)]
    H -->|Termasuk Link Gambar| E1

    F --> F1[Transaksi Stok Masuk]
    F --> F2[Transaksi Stok Keluar]
    F1 --> I{Periksa Stok}
    F2 --> I
    I -->|Stok Tersedia| J[Perbarui Stok]
    I -->|Stok Tidak Cukup| K[Tolak Transaksi]
    J --> L[(Database: Transaksi)]
    K --> M[Respon Error]
    L --> N[Daftar Riwayat Transaksi]

    A:::start
    B:::process
    C:::decision
    D:::process
    E:::process
    F:::process
    E1:::process
    F1:::process
    F2:::process
    I:::decision
    J:::process
    K:::error
    M:::error
    N:::process
    H:::data
    L:::data

    classDef start shape:circle
    classDef process shape:rect
    classDef decision shape:diamond
    classDef data shape:cylinder
    classDef error shape:rect


```

## ERD Diagram

```mermaid
erDiagram
    User ||--o{ Transactions : makes
    Categories ||--o{ Products : contains
    Products ||--o{ Transactions : involved_in
    Categories ||--o{ Products_Category : categorized_by
    Products ||--o{ Products_Category : has

    User {
        int users_id PK
        string name
        string email
        string password
        string role
    }

    Categories {
        int category_id PK
        string name
        string description
    }

    Products {
        int product_id PK
        string name
        string image_url
        decimal purchase_price
        decimal selling_price
        int stock
    }

    Products_Category {
        int product_id FK
        int category_id FK
    }

    Transactions {
        int transaction_id PK
        int product_id FK
        int users_id FK
        string type
        int quantity
        decimal total_price
        timestamp transaction_date
    }

```

## User Requirement

| ID                      | UR-01                                                                                                                                                   |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Judul**               | Manajemen Produk                                                                                                                                        |
| **Aktor**               | Admin                                                                                                                                                   |
| **Deskripsi**           | Admin dapat menambahkan, melihat, mengedit, dan menghapus data produk yang mencakup informasi nama, harga beli, harga jual, stok, gambar, dan kategori. |
| **Tujuan**              | Memudahkan admin dalam mengelola data produk yang tersedia di sistem, sehingga data selalu up-to-date dan akurat.                                       |
| **Kriteria Penerimaan** | - Admin dapat melihat seluruh daftar produk lengkap dengan kategori.                                                                                    |

| ID                      | UR-02                                                                                                        |
| ----------------------- | ------------------------------------------------------------------------------------------------------------ |
|                         |                                                                                                              |
| **Judul**               | Admin Melakukan Transaksi dan Melihat Riwayat Transaksi                                                      |
| **Aktor**               | Admin                                                                                                        |
| **Deskripsi**           | Admin dapat membuat transaksi pembelian produk dan melihat daftar riwayat transaksi.                         |
| **Tujuan**              | Memungkinkan admin mencatat pembelian produk ke sistem dan mengecek histori transaksi yang pernah dilakukan. |
| **Kriteria Penerimaan** | - Admin dapat memilih produk dan jumlahnya untuk membuat transaksi baru.                                     |
