CREATE DATABASE IF NOT EXISTS mbg_smart_logistics;
USE mbg_smart_logistics;

CREATE TABLE pengguna (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    kata_sandi VARCHAR(255) NOT NULL,
    peran ENUM('admin', 'kurir', 'guru') NOT NULL,
    tipe_kendaraan ENUM('motor', 'mobil', 'tidak_ada') DEFAULT 'tidak_ada',
    kapasitas_kendaraan INT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
);

CREATE TABLE lokasi (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nama_lokasi VARCHAR(150) NOT NULL,
    tipe_lokasi ENUM('SPPG', 'SEKOLAH', 'ESTAFET') NOT NULL,
    latitude DOUBLE NOT NULL,
    longitude DOUBLE NOT NULL,
    kebutuhan_porsi INT DEFAULT 0,
    kontak_pic VARCHAR(100) NULL,
    batas_waktu VARCHAR(10) NULL,
    waktu_layanan_menit INT DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
);

CREATE TABLE rute (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    kurir_id BIGINT NOT NULL,
    tanggal DATE NOT NULL,
    total_jarak_km DOUBLE DEFAULT 0,
    waktu_mulai_aktual TIME NULL,
    waktu_selesai_aktual TIME NULL,
    status_rute ENUM('TUNDA', 'PROSES', 'SELESAI') DEFAULT 'TUNDA',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    CONSTRAINT fk_rute_kurir FOREIGN KEY (kurir_id) REFERENCES pengguna(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE pemberhentian_rute (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    rute_id BIGINT NOT NULL,
    lokasi_id BIGINT NOT NULL,
    urutan_berhenti INT NOT NULL,
    waktu_tiba_aktual DATETIME NULL,
    porsi_turun INT DEFAULT 0,
    porsi_naik INT DEFAULT 0,
    bukti_foto VARCHAR(255) NULL,
    catatan TEXT NULL,
    status_perhentian ENUM('MENUNGGU', 'TIBA', 'SELESAI', 'GAGAL') DEFAULT 'MENUNGGU',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    CONSTRAINT fk_pemberhentian_rute FOREIGN KEY (rute_id) REFERENCES rute(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_pemberhentian_lokasi FOREIGN KEY (lokasi_id) REFERENCES lokasi(id) ON DELETE CASCADE ON UPDATE CASCADE
);