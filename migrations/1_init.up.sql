CREATE TABLE users (
    id CHAR(36) PRIMARY KEY, -- UUID
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id CHAR(36) PRIMARY KEY, -- UUID
    name VARCHAR(255) NOT NULL UNIQUE,  -- Food/Drink/Sport/Pet/Transport/Necessities/Medical/Education
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE category_images (
    id CHAR(36) PRIMARY KEY, -- UUID
    image_url VARCHAR(255) NOT NULL UNIQUE,
    category_id CHAR(36) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE user_categories (
    id CHAR(36) PRIMARY KEY, -- UUID
    user_id CHAR(36) NOT NULL,
    category_image_id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type ENUM('Income', 'Expense') NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (category_image_id) REFERENCES category_images(id) ON DELETE CASCADE
);

CREATE TABLE transactions (
    id CHAR(36) PRIMARY KEY, -- UUID
    user_id CHAR(36) NOT NULL,
    category_image_url VARCHAR(255) NOT NULL,
    category_name VARCHAR(255) NOT NULL,
    description VARCHAR(255) DEFAULT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    transaction_date DATE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
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


-- Insert Default Category
INSERT INTO categories (id, name) VALUES ('53eb757c-49aa-4cb0-af88-60d4402ad051', "Food");
INSERT INTO categories (id, name) VALUES ('2b214a04-02f7-49a0-a7fe-2b3fd64e713a', "Drink");
INSERT INTO categories (id, name) VALUES ('b787a979-ba5c-433d-a035-52f1f3b1f9a5', "Jeans");
INSERT INTO categories (id, name) VALUES ('002937e4-0219-4869-ba77-77b1da1b6eb4', "Medical");

INSERT INTO category_images (id, image_url, category_id) VALUES ('64918708-021e-4dbb-8d12-c326e3231ba4', "./category_images/food.png", '53eb757c-49aa-4cb0-af88-60d4402ad051');
INSERT INTO category_images (id, image_url, category_id) VALUES ('45528ca6-9984-446e-bb3f-1abb9eefea22', "./category_images/dango.png", '53eb757c-49aa-4cb0-af88-60d4402ad051');
INSERT INTO category_images (id, image_url, category_id) VALUES ('ac599cc5-b565-4f79-8cd6-05ea921a10fb', "./category_images/drink.png", '2b214a04-02f7-49a0-a7fe-2b3fd64e713a');
INSERT INTO category_images (id, image_url, category_id) VALUES ('4f47676a-4eae-4aa2-bd77-04bcf0d6d416', "./category_images/jeans.png", 'b787a979-ba5c-433d-a035-52f1f3b1f9a5');
INSERT INTO category_images (id, image_url, category_id) VALUES ('d22d4b1c-4288-499d-bf5e-9c6dc2d4dbb4', "./category_images/medical.png", '002937e4-0219-4869-ba77-77b1da1b6eb4');

-- Default User Category
-- INSERT INTO user_categories (id, user_id, category_image_id, name, type) VALUES ('48ed5950-8fa9-4f14-beda-e474af919220', null, '64918708-021e-4dbb-8d12-c326e3231ba4', "Food", "Expense");
-- INSERT INTO user_categories (id, user_id, category_image_id, name, type) VALUES ('eaaec2c6-5b66-4248-9693-53c18d700e15', null, '64918708-021e-4dbb-8d12-c326e3231ba4', "Fast Food", "Expense");
-- INSERT INTO user_categories (id, user_id, category_image_id, name, type) VALUES ('2c1be799-749b-48d0-863a-9931c82d167e', null, '45528ca6-9984-446e-bb3f-1abb9eefea22', "Snack", "Expense");
-- INSERT INTO user_categories (id, user_id, category_image_id, name, type) VALUES ('28b508b0-ba1b-469a-956b-b7edaeac85ff', null, 'ac599cc5-b565-4f79-8cd6-05ea921a10fb', "Drink", "Expense");
-- INSERT INTO user_categories (id, user_id, category_image_id, name, type) VALUES ('ce12b949-2957-4d61-9bc7-5014c0ee84f7', null, '4f47676a-4eae-4aa2-bd77-04bcf0d6d416', "Jeans", "Expense");
-- INSERT INTO user_categories (id, user_id, category_image_id, name, type) VALUES ('91013937-0abe-4418-9dc0-e480e6f51d1b', null, 'd22d4b1c-4288-499d-bf5e-9c6dc2d4dbb4', "Medicine", "Expense");

-- Testing
-- INSERT INTO transactions (id, user_id, category_id, amount, transaction_date, description) VALUES ('858da5cc-d58d-42df-8dde-c0cecd8fac90', '00fc454f-83ea-44a2-845d-b80f5a9b6f3a', '48ed5950-8fa9-4f14-beda-e474af919220', 100.00, '2017-06-15', 'pizza');

INSERT INTO users (id, username, email, password_hash) VALUES ('00fc454f-83ea-44a2-845d-b80f5a9b6f3a', 'test', "test@gmail.com", "test");
