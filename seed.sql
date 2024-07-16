INSERT INTO roles(id,name) VALUES (1,'lecturer'),(2,'student');

INSERT INTO users(id,name, email, password, role_id) VALUES
(1,'Paul Mwaniki','paulmwaniki@gmail.com','qwerty1234',1),
(2,'Matthew Thiongo','matthewthiongo@gmail.com','qwerty1234',1),
(3,'Florence Kimani','florencekimani@gmail.com','qwerty1234',1),
(4,'David Kirop','davidkirop@gmail.com','qwerty1234',1),

(5,'Jordan Peterson','jordanpeterson@gmail.com','1234567',2),
(6,'Max Verstapen','maxverstapen@gmail.com','1234567',2),
(7,'Oscar Piastri','oscarpiastri@gmail.com','1234567',2),
(8,'Lando Norris','landonorris@gmail.com','1234567',2),
(9,'Carlos Sainz','carlossainz@gmail.com','1234567',2)
;

INSERT INTO projects(id,student_id,lecturer_id,title,description,start_date,end_date) VALUES
(1,5,1,'Social network','building a social platform that is exclusive to the school only.Will improve communication in school between students','2024-05-15T14:48:00.000Z','2024-08-15T14:48:00.000Z'),
(2,6,1,'Leave Management system','A system that will allow company members to ask for leave and streamline the process of knowing if its the right time or not','2024-05-15T14:48:00.000Z','2024-08-15T14:48:00.000Z'),
(3,7,1,'Software jobs and internships','An app that helps college and university graduates find internship openings and also companies that are hiring Be specific to the Kenyan market Also can add tech meetups that are happening to help with networking','2024-05-15T14:48:00.000Z','2024-08-15T14:48:00.000Z'),
(4,8,2,'Chama Money Tracker','Having the different and many ways people can borrow money and different types of instalments it can be hectic and sometimes hard to figure out what is going on','2024-05-15T14:48:00.000Z','2024-08-15T14:48:00.000Z'),
(5,9,2,'QR Code Payment','Leveraging the power of a QR code and smartphone camera where once you reach your destination with the boda boda you scan the QR code and prompts you to pay','2024-05-15T14:48:00.000Z','2024-08-15T14:48:00.000Z');


INSERT INTO submissions(id,projec_id,student_id,lecturer_id,submission_date,document_path,description,reviewed) VALUES 
(1,1,5,1,'2024-05-30T14:48:00.000Z','https://googledrive/path/to/document','finished chapter1 added the reference that you asked for and made it to times new roman font',false),
(2,2,6,1,'2024-05-30T14:48:00.000Z','https://googledrive/path/to/document','models for the database have been updated and being used in production .images of the current version are on the last page ',false),
(3,3,7,1,'2024-05-30T14:48:00.000Z','https://googledrive/path/to/document','Finished my chapter 2 which was literature review able to find references about the design i want to impelement ',false);

INSERT INTO feedbacks(id,submission_id,lecturer_id,feedback_date,comments) VALUES 
()