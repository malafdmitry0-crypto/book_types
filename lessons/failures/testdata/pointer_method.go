package invalid
type Worker interface{Work()};type Job struct{};func(*Job)Work(){};var _ Worker=Job{}
