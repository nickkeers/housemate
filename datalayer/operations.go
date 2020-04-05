package datalayer

type HousemateDataOperations interface {
    // HouseholdMember operations
    GetHouseholdMemberById(userId int) (*HouseholdMember, error)
    LookupHouseholdMember(email string) (*HouseholdMember, error)
    DeleteHouseholdMember(userId int) error
    AddNewHouseholdMember(member *HouseholdMember) error
}
