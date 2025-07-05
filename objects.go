package main

type Opportunity struct {
	oppId      int
	oppNumber  int
	oppName    string
	oppOwner   Contact
	oppLead    Lead
	oppAccount Account
}

type Lead struct {
	leadId      int
	leadNumber  int
	leadOwner   Contact
	leadAccount Account
	leadSource  string
}

type Contact struct {
	contactId         int
	contactSurname    string
	contactName       string
	contactMiddlename string
	contactCompany    *Account
	contactPhone      string
	contactEmail      string
	contactJobTitle   string
	contactOwner      *Contact
}

type Account struct {
	accountId          int
	accountName        string
	accountFullname    string
	accountMainContact Contact
	accountOwner       Contact
	accountInn         uint64
	accountWeb         string
}
