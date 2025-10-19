CREATE TABLE IF NOT EXISTS order_items(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `order_id` INT UNSIGNED NOT NULL,
    `product_id` INT UNSIGNED NOT NULL,
    `quantity` INT NOT NULL,
    `price` DECIMAL(10,2) NOT NULL,
    FOREIGN KEY(`order_id`) REFERENCES orders(`id`) ON DELETE CASCADE,
    -- When you delete a record from a parent table, ON DELETE CASCADE
    -- automatically deletes all related records in the child table 
    -- that reference the deleted parent record through a foreign key relationship.
    FOREIGN KEY(`product_id`) REFERENCES products(`id`) ON DELETE CASCADE
);