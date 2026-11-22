# Audit sumber Mika untuk migrasi ke Mikaoka

Tanggal pemeriksaan: 2026-11-22

## Repositori sumber
- Upstream asli: https://github.com/wangzan101/Mika
- Fork pengguna: https://github.com/project10dollar-png/Mika
- Target: https://github.com/project10dollar-png/Mikaoka

## Temuan repositori upstream
- Default branch: `main`
- Commit yang diaudit: `2ca391c4c868a39d98d06639f327489baaebade8`
- Ukuran GitHub: sekitar 322,004 KB.
- Inventaris Git tree: 4,685 entri; 4,091 file dan 594 direktori.
- Ukuran konten file yang terdaftar di tree: sekitar 109,011,996 byte.
- Tree recursive tidak terpotong.

## Temuan fork dan perubahan HEY OKA
- Fork `project10dollar-png/Mika` tersedia dan dapat ditulis oleh pemilik akun.
- Branch yang terdeteksi: `main` dan `feature/hey-oka-branding`.
- Branch fitur HEY OKA harus dipertahankan sebagai sumber perubahan pengguna.
- Jangan menimpa perubahan HEY OKA dengan upstream `main` tanpa menggabungkan/mengevaluasi perbedaannya.

## Gambaran arsitektur yang dibaca
- Proyek Android berbasis Kotlin dan Gradle.
- Modul Gradle: `:app` dan `:core`.
- Android namespace/application ID: `com.wangzan101.mika`.
- Minimum Android SDK 29, target SDK 35, compile SDK 36.
- Java/Kotlin toolchain 17.
- Fitur mencakup kendali aplikasi Android, browser, model AI/provider, percakapan, penjadwalan, file/sandbox, serta CLI dan relay.
- CI menjalankan tes JVM untuk core, tes CLI Python, tes/build Android, dan tes relay Go.
- Runtime sandbox besar disimpan di `app/src/main/assets/runtime/rootfs.tar.xz` (sekitar 59 MB); file besar lain termasuk demo media dan aset gambar.

## Berkas yang diperiksa secara langsung
- `README.md`
- `settings.gradle.kts`
- `build.gradle.kts`
- `app/build.gradle.kts`
- `.github/workflows/ci.yml`

## Status migrasi
Dokumen ini adalah hasil inventaris dan pembacaan awal, bukan salinan seluruh kode.
Belum ada bukti bahwa 4,091 file upstream dan seluruh aset biner telah tersalin ke Mikaoka.
Salinan lengkap memerlukan operasi Git clone/mirror atau transfer Git yang mempertahankan file biner, tree, dan riwayat. GitHub API biasa tidak dapat menggunakan blob SHA lintas-repositori secara langsung.
Untuk mendorong berkas workflow dari workflow otomatis, kredensial harus memiliki izin workflow selain izin contents.
Jangan menyimpan token akses di repository atau file dokumentasi.

## Referensi
- Upstream tree: https://github.com/wangzan101/Mika/tree/main
- Fork: https://github.com/project10dollar-png/Mika
- HEY OKA branch: https://github.com/project10dollar-png/Mika/tree/feature/hey-oka-branding
