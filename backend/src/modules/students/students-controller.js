const asyncHandler = require("express-async-handler");
const {
    getAllStudents,
    addNewStudent,
    getStudentDetail,
    setStudentStatus,
    updateStudent
} = require("./students-service");


const handleGetAllStudents = asyncHandler(async (req, res) => {
    const filters = req.query;
    const students = await getAllStudents(filters);
    res.status(200).json(students);
});

const handleAddStudent = asyncHandler(async (req, res) => {
    const payload = req.body;
    const result = await addNewStudent(payload);
    res.status(201).json(result);
});

const handleUpdateStudent = asyncHandler(async (req, res) => {
    const id = parseInt(req.params.id);
    const payload = {
        ...req.body,
        userId: id,
    };
    const result = await updateStudent(payload);
    res.status(200).json(result);
});


const handleGetStudentDetail = asyncHandler(async (req, res) => {
    const id = parseInt(req.params.id);
    const student = await getStudentDetail(id);
    res.status(200).json(student);
});

const handleStudentStatus = asyncHandler(async (req, res) => {
    const userId = parseInt(req.params.id);
    const { reviewerId, status } = req.body;
    const result = await setStudentStatus({ userId, reviewerId, status });
    res.status(200).json(result);
});

module.exports = {
    handleGetAllStudents,
    handleGetStudentDetail,
    handleAddStudent,
    handleStudentStatus,
    handleUpdateStudent,
};
