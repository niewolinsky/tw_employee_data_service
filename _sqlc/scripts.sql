-- DROP SCRIPT
DROP TABLE IF EXISTS public.performance CASCADE;
DROP TABLE IF EXISTS public.zone CASCADE;
DROP TABLE IF EXISTS public.artist CASCADE;
DROP TABLE IF EXISTS public.happening CASCADE;
DROP TABLE IF EXISTS public.place CASCADE;
DROP TYPE IF EXISTS ENUM_SUBSCRIPTION_STATUS;
DROP TYPE IF EXISTS ENUM_VERIFICATION_STATUS;

-- INSERT SCRIPT
INSERT INTO public.place
(name, description, type, longitude, latitude, verification_status, subscription_status, phone, email, city, country, website, opening_hours)
VALUES
('Club A', 'Best club in town A', 'Club', 12.345678, 23.456789, 'PUBLISHED', 'ON', '+123456789', 'contact@clubA.com', 'CityA', 'CountryA', 'www.clubA.com', '{"monday": "10:00-23:00"}'),
('Club B', 'Best club in town B', 'Club', 24.567890, 34.678901, 'PUBLISHED', 'ON', '+234567890', 'contact@clubB.com', 'CityB', 'CountryB', 'www.clubB.com', '{"tuesday": "12:00-24:00"}'),
('Club C', 'Best club in town C', 'Club', 35.789012, 45.890123, 'HIDDEN', 'OFF', '+345678901', 'contact@clubC.com', 'CityC', 'CountryC', 'www.clubC.com', '{"wednesday": "14:00-02:00"}');

INSERT INTO public.happening
(name, description, longitude, latitude, verification_status, place_id, city, country, start_time, end_time)
VALUES
('Fest A', 'Top Fest A', 45.890124, 56.901234, 'PUBLISHED', 1, 'CityD', 'CountryD', '2023-10-12 10:00:00', '2023-10-13 10:00:00'),
('Fest B', 'Top Fest B', 56.901235, 67.012345, 'HIDDEN', 2, 'CityE', 'CountryE', '2023-11-15 12:00:00', '2023-11-16 12:00:00'),
('Fest C', 'Top Fest C', 67.012346, 78.123456, 'SCRAPPED', NULL, 'CityF', 'CountryF', '2023-12-18 14:00:00', '2023-12-19 14:00:00');

INSERT INTO public.zone
(place_id, happening_id, name)
VALUES
(1, NULL, 'Zone 1A'),
(2, NULL, 'Zone 2B'),
(NULL, 1, 'Zone 3C');

INSERT INTO public.artist
(name)
VALUES
('Artist A'),
('Artist B'),
('Artist C');

INSERT INTO public.performance
(zone_id, artist_id, artist_name, start_time, end_time)
VALUES
(1, 1, 'Artist A', '2023-10-12 11:00:00', '2023-10-12 14:00:00'),
(2, 2, 'Artist B', '2023-10-13 15:00:00', '2023-10-13 18:00:00'),
(3, 3, 'Artist C', '2023-10-14 19:00:00', '2023-10-14 22:00:00');
