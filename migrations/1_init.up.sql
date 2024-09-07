CREATE TABLE user (
    user_id CHAR(36) PRIMARY KEY, -- UUID
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE category_image (
    image_id CHAR(36) PRIMARY KEY, -- UUID
    url VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE category (
    category_id CHAR(36) PRIMARY KEY, -- UUID
    user_id CHAR(36) DEFAULT NULL,
    image_id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type ENUM('Income', 'Expense') NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (image_id) REFERENCES category_image(image_id)
);

CREATE TABLE transaction (
    transaction_id CHAR(36) PRIMARY KEY, -- UUID
    user_id CHAR(36) NOT NULL,
    category_id CHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    transaction_date DATE NOT NULL,
    description VARCHAR(255) DEFAULT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES category(category_id)
);

-- CREATE TABLE Budget (
--     budget_id CHAR(36) PRIMARY KEY, -- UUID
--     user_id CHAR(36),
--     category_id CHAR(36),
--     amount DECIMAL(10, 2) NOT NULL,
--     start_date DATE NOT NULL,
--     end_date DATE NOT NULL,
--     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
--     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
--     FOREIGN KEY (category_id) REFERENCES Category(category_id) ON DELETE CASCADE
-- );

-- CREATE TABLE Savings (
--     savings_id CHAR(36) PRIMARY KEY, -- UUID
--     user_id CHAR(36),
--     goal_name VARCHAR(255) NOT NULL,
--     target_amount DECIMAL(10, 2) NOT NULL,
--     current_amount DECIMAL(10, 2) NOT NULL,
--     target_date DATE NOT NULL,
--     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
--     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
-- );

INSERT INTO user (user_id, username, email, password_hash) VALUES ('00fc454f-83ea-44a2-845d-b80f5a9b6f3a', 'test', "test@gmail.com", "test");

INSERT INTO category_image (image_id, url) VALUES ('64918708-021e-4dbb-8d12-c326e3231ba4', "../personal-financial-tracker-frontend/food.png");

INSERT INTO category (category_id, user_id, image_id, name, type) VALUES ('48ed5950-8fa9-4f14-beda-e474af919220', null, '64918708-021e-4dbb-8d12-c326e3231ba4', "Food", "Expense");

INSERT INTO transaction (transaction_id, user_id, category_id, amount, transaction_date, description) VALUES ('858da5cc-d58d-42df-8dde-c0cecd8fac90', '00fc454f-83ea-44a2-845d-b80f5a9b6f3a', '48ed5950-8fa9-4f14-beda-e474af919220', 100.00, '2017-06-15', 'pizza');
