INSERT INTO articles (title, text, created_by_id)
VALUES ('What happened', 'wth', 1),
    ('War', 'No war pls', 2),
    ('Herba life', 'Drink it', 1),
    ('O ZOne', 'Some song', 12),
    ('Destiny', 'Another song', 10),
    ('SS', 'In my life', 14),
    ('Here', 'Random text there', 3)
RETURNING *;