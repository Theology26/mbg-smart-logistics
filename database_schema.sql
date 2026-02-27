CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role ENUM('vendor', 'guru') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE ingredient_qcs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    ingredient_name VARCHAR(100) NOT NULL,
    photo_url VARCHAR(255) NOT NULL,
    status ENUM('SAFE', 'UNSAFE') NOT NULL,
    ai_analysis TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE delivery_plans (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    menu_name VARCHAR(150) NOT NULL,
    school_name VARCHAR(150) NOT NULL,
    distance_km FLOAT NOT NULL,
    travel_time_hr FLOAT NOT NULL,
    status ENUM('APPROVED', 'REJECTED') NOT NULL,
    ai_reason TEXT,
    delivery_date DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE confirmations ( 
    id INT AUTO_INCREMENT PRIMARY KEY, 
    delivery_plan_id INT NOT NULL UNIQUE, 
    guru_id INT NOT NULL, 
    photo_proof VARCHAR(255), 
    received_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (delivery_plan_id) REFERENCES delivery_plans(id) ON DELETE CASCADE, 
    FOREIGN KEY (guru_id) REFERENCES users(id) ON DELETE CASCADE 
) ENGINE=InnoDB;

CREATE TABLE feedbacks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    delivery_plan_id INT NOT NULL UNIQUE,
    rating INT NOT NULL,
    comment TEXT,
    is_fresh BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (delivery_plan_id) REFERENCES delivery_plans(id) ON DELETE CASCADE
) ENGINE=InnoDB;