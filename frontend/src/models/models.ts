export interface Quiz {
    id: string;
    name: string;
    questions: QuizQuestion[];
}

export interface QuizQuestion {
    id: string;
    name: string;
    quizChoices: QuizChoice[];
}

export interface QuizChoice {
    id: string;
    name: string;
    correct: boolean;
}