CREATE TABLE IF NOT EXISTS owners (
    `id` BINARY(16),
    `fullName` VARCHAR(100),
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS pets (
    `id` BINARY(16),
    `name` VARCHAR(100),
    `breed` VARCHAR(100),
    `ownerId` BINARY(16),
    PRIMARY KEY (`id`),
    CONSTRAINT `FK_pets_owner` FOREIGN KEY (`ownerId`) REFERENCES owners(`id`) ON DELETE CASCADE
)