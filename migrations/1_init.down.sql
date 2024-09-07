CREATE TABLE User (
    user_id CHAR(36) PRIMARY KEY, -- UUID
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE Category_Image (
    image_id CHAR(36) PRIMARY KEY, -- UUID
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE Category (
    category_id CHAR(36) PRIMARY KEY, -- UUID
    user_id CHAR(36) DEFAULT NULL,
    image_id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type ENUM('Income', 'Expense') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES User(user_id) ON DELETE CASCADE,
    FOREIGN KEY (image_id) REFERENCES Category_Image(image_id)
);

CREATE TABLE Transaction (
    transaction_id CHAR(36) PRIMARY KEY, -- UUID
    user_id CHAR(36) NOT NULL,
    category_id CHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    transaction_date DATE NOT NULL,
    description VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES User(user_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES Category(category_id)
);

-- CREATE TABLE Budget (
--     budget_id CHAR(36) PRIMARY KEY, -- UUID
--     user_id CHAR(36),
--     category_id CHAR(36),
--     amount DECIMAL(10, 2) NOT NULL,
--     start_date DATE NOT NULL,
--     end_date DATE NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE,
--     FOREIGN KEY (category_id) REFERENCES Categories(category_id) ON DELETE CASCADE
-- );

-- CREATE TABLE Savings (
--     savings_id CHAR(36) PRIMARY KEY, -- UUID
--     user_id CHAR(36),
--     goal_name VARCHAR(255) NOT NULL,
--     target_amount DECIMAL(10, 2) NOT NULL,
--     current_amount DECIMAL(10, 2) NOT NULL,
--     target_date DATE NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
-- );
