SELECT
  feedbacks.id AS feedback_id,
  feedbacks.feedback_date,
  feedbacks.comments,
  submissions.id AS submission_id,
  submissions.submission_date,
  submissions.document_path,
  submissions.description,
  users.name AS student_name,
  users.email AS student_email
FROM
  feedbacks
  JOIN submissions ON feedbacks.submission_id = submissions.id
  JOIN users ON submissions.student_id = users.id;