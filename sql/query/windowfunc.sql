-- WINDOW FUNCTION - RANKING MEMBER BERDASARKAN POIN
SELECT 
  id,
  name,
  reward_points,
  RANK() OVER (ORDER BY reward_points DESC) AS rank
FROM members;