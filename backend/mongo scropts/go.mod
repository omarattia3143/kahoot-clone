// MongoDB Playground
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.

// The current database to use.
use("quiz");

db.quizzes.insert({"title": "Quiz 1", "questions": [{"question": "What is 1+1?", "answers": ["1", "2", "3", "4"], "correctAnswer": "2"}]});