package routes

import (
	"covi-helper-backend/structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

type UserController struct {
	session *sql.DB
}

//var wg sync.WaitGroup {will need to call free(wg) for destroying this global attribute}

func NewUserController(s *sql.DB) *UserController {
	return &UserController{s}
}

func (uc UserController) Location(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var _ = ""
	var latitude = 30.6684076
	var longitude = 76.8627135
	var _ string
	url := "https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=" + fmt.Sprint(latitude) + "&longitude=" + fmt.Sprint(longitude) + "&localityLanguage=en"
	//url := "https://maps.googleapis.com/maps/api/geocode/json?key=" + apikey + "&latlng=" + fmt.Sprint(latitude) + "," + fmt.Sprint(longitude)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}

func (uc UserController) GetStates(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var u []*structs.States
	sqlStatement, err := uc.session.Query(" SELECT * from masterdb.states WHERE country_id=?;", 101)
	if err != nil {
		panic(err)
	}
	defer sqlStatement.Close()
	for sqlStatement.Next() {
		c := new(structs.States) // initialize a new instance
		error := sqlStatement.Scan(&c.Id, &c.Name, &c.CountryId)
		if error != nil {
			fmt.Println(error)
		}
		u = append(u, c) // add each instance to the slice
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) RegisterVolunteer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	u := structs.Volunteer{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Allowed = false
	sqlStatement, err := uc.session.Prepare("INSERT IGNORE INTO masterdb.loggedin (email_id, user_id, name, picture, phone_number) VALUES (?, ?, ?, ?,?);")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.EmailId, u.UserId, u.Name, u.Picture, u.PhoneNumber)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound) // 302
	}
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		fmt.Println(err)
	}
}

func (uc UserController) AddOxygen(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("INSERT INTO masterdb.oxygen_table (name,lead_type, address, map_link, phone_number, whatsapp_number, website, notes, creation_time, statecode, citycode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Website, u.Notes, u.CreationTime, u.Statecode, u.Citycode)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) AddBed(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("INSERT INTO masterdb.beds_table (name,lead_type, address, map_link, phone_number, whatsapp_number, notes, creation_time, statecode, citycode, website) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Notes, u.CreationTime, u.Statecode, u.Citycode, u.Website)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) AddPlasma(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("INSERT INTO masterdb.plasma_table (name,lead_type, address, map_link, phone_number, whatsapp_number, notes, creation_time, statecode, citycode, website) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Notes, u.CreationTime, u.Statecode, u.Citycode, u.Website)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) AddMedicine(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("INSERT INTO masterdb.medicine_table (name,lead_type, address, map_link, phone_number, whatsapp_number, notes, creation_time, statecode, citycode, website) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Notes, u.CreationTime, u.Statecode, u.Citycode, u.Website)
	if err != nil {
		fmt.Println(err)
		return
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) AddAmbulance(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("INSERT INTO masterdb.ambulance_table (name,lead_type, address, map_link, phone_number, whatsapp_number, notes, creation_time, statecode, citycode, website) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Notes, u.CreationTime, u.Statecode, u.Citycode, u.Website)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		fmt.Println(err)
	}
}

func (uc UserController) AddHelpingHand(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("INSERT INTO masterdb.helping_table (name,lead_type, address, map_link, phone_number, whatsapp_number, notes, creation_time, statecode, citycode, website) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Notes, u.CreationTime, u.Statecode, u.Citycode, u.Website)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) GetResources(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	cityCode := p.ByName("cityCode")
	stateName := p.ByName("stateName")
	cityName := p.ByName("cityName")
	baseStruct := new(structs.ResourceMaster)
	ack := make(chan bool, 6)
	go func() {
		sqlStatement, err := uc.session.Query("SELECT * FROM masterdb.ambulance_table WHERE citycode=?", cityCode)
		if err != nil {
			fmt.Println(err)
		}
		defer sqlStatement.Close()
		for sqlStatement.Next() {
			c := new(structs.ResourceLead)
			err = sqlStatement.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
			if err != nil {
				fmt.Println(err)
			}
			baseStruct.Ambulance = append(baseStruct.Ambulance, c)
		}
		ack <- true
	}()
	go func() {
		query, err := uc.session.Query("SELECT * FROM masterdb.oxygen_table WHERE citycode=?", cityCode)
		if err != nil {
			fmt.Println(err)
		}
		defer query.Close()
		for query.Next() {
			c := new(structs.ResourceLead)
			err = query.Scan(&c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime, &c.Id, &c.Website)
			if err != nil {
				fmt.Println(err)
			}
			baseStruct.Oxygen = append(baseStruct.Oxygen, c)
		}
		ack <- true
	}()
	go func() {
		query, err := uc.session.Query("SELECT * FROM masterdb.plasma_table WHERE citycode=?", cityCode)
		if err != nil {
			fmt.Println(err)
		}
		defer query.Close()
		for query.Next() {
			c := new(structs.ResourceLead)
			err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
			if err != nil {
				fmt.Println(err)
			}
			baseStruct.Plasma = append(baseStruct.Plasma, c)
		}
		ack <- true
	}()
	go func() {
		query, err := uc.session.Query("SELECT * FROM masterdb.medicine_table WHERE citycode=?", cityCode)
		if err != nil {
			fmt.Println(err)
		}
		defer query.Close()
		for query.Next() {
			c := new(structs.ResourceLead)
			err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
			if err != nil {
				fmt.Println(err)
			}
			baseStruct.Medicine = append(baseStruct.Medicine, c)
		}
		ack <- true
	}()
	go func() {
		query, err := uc.session.Query("SELECT * FROM masterdb.helping_table WHERE citycode=?", cityCode)
		if err != nil {
			fmt.Println(err)
		}
		defer query.Close()
		for query.Next() {
			c := new(structs.ResourceLead)
			err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
			if err != nil {
				fmt.Println(err)
			}
			baseStruct.HelpingHand = append(baseStruct.HelpingHand, c)
		}
		ack <- true
	}()
	go func() {
		query, err := uc.session.Query("SELECT * FROM masterdb.beds_table WHERE citycode=?", cityCode)
		if err != nil {
			fmt.Println(err)
		}
		defer query.Close()
		for query.Next() {
			c := new(structs.ResourceLead)
			err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
			if err != nil {
				fmt.Println(err)
			}
			baseStruct.Bed = append(baseStruct.Bed, c)
		}
		ack <- true
	}()
	for i := 0; i < 6; i++ {
		fmt.Println(<-ack)
	}
	masterResponse := new(structs.ResourceMain)
	masterResponse.ResourceMaster = append(masterResponse.ResourceMaster, baseStruct)
	masterResponse.ResponseCode = 1
	masterResponse.CityCode = cityCode
	masterResponse.StateName = stateName
	masterResponse.CityName = cityName
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	err := json.NewEncoder(w).Encode(masterResponse)
	if err != nil {
		fmt.Println(err)
	}
}

func (uc UserController) VerifyResources(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	emailId := p.ByName("emailId")
	var allowed bool
	baseStruct := new(structs.ResourceMaster)
	ack := make(chan bool, 6)
	queryVolunteers, er := uc.session.Query("SELECT allowed FROM masterdb.loggedin WHERE email_id=?", emailId)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer queryVolunteers.Close()
	for queryVolunteers.Next() {
		err := queryVolunteers.Scan(&allowed)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if allowed != false {

		go func() {
			sqlStatement, err := uc.session.Query("SELECT * FROM masterdb.ambulance_table WHERE verified_by=?;", "")
			if err != nil {
				fmt.Println(err)
			}
			defer sqlStatement.Close()
			for sqlStatement.Next() {
				c := new(structs.ResourceLead)
				err = sqlStatement.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Ambulance = append(baseStruct.Ambulance, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.oxygen_table WHERE verified_by=?;", "")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime, &c.Id, &c.Website)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Oxygen = append(baseStruct.Oxygen, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.plasma_table WHERE verified_by=?;", "")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Plasma = append(baseStruct.Plasma, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.medicine_table WHERE verified_by=?;", "")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Medicine = append(baseStruct.Medicine, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.helping_table WHERE verified_by=?;", "")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.HelpingHand = append(baseStruct.HelpingHand, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.beds_table WHERE verified_by=?;", "")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Bed = append(baseStruct.Bed, c)
			}
			ack <- true
		}()
		for i := 0; i < 6; i++ {
			fmt.Println(<-ack)
		}
		verifyResource := new(structs.VerifyResource)
		verifyResource.ResourceMaster = append(verifyResource.ResourceMaster, baseStruct)
		verifyResource.ResponseCode = 1
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		err := json.NewEncoder(w).Encode(verifyResource)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		verifyResource := new(structs.VerifyResource)
		verifyResource.ResponseCode = 0
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound) // 302
		err := json.NewEncoder(w).Encode(verifyResource)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (uc UserController) VerifyOxygenLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("UPDATE masterdb.oxygen_table SET name=?, lead_type=?, address=?, map_link=?, phone_number=?, whatsapp_number=?, website=?, notes=?, verified_time=?, verified_by=?, statecode=?, citycode=? WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Website, u.Notes, u.VerfiedTime, u.VerifiedBy, u.Statecode, u.Citycode, u.Id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) VerifyBedLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("UPDATE masterdb.beds_table SET name=?, lead_type=?, address=?, map_link=?, phone_number=?, whatsapp_number=?, website=?, notes=?, verified_time=?, verified_by=?, statecode=?, citycode=? WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Website, u.Notes, u.VerfiedTime, u.VerifiedBy, u.Statecode, u.Citycode, u.Id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) VerifyPlasmaLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("UPDATE masterdb.plasma_table SET name=?, lead_type=?, address=?, map_link=?, phone_number=?, whatsapp_number=?, website=?, notes=?, verified_time=?, verified_by=?, statecode=?, citycode=? WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Website, u.Notes, u.VerfiedTime, u.VerifiedBy, u.Statecode, u.Citycode, u.Id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) VerifyMedicineLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("UPDATE masterdb.medicine_table SET name=?, lead_type=?, address=?, map_link=?, phone_number=?, whatsapp_number=?, website=?, notes=?, verified_time=?, verified_by=?, statecode=?, citycode=? WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Website, u.Notes, u.VerfiedTime, u.VerifiedBy, u.Statecode, u.Citycode, u.Id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) VerifyAmbulanceLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("UPDATE masterdb.ambulance_table SET name=?, lead_type=?, address=?, map_link=?, phone_number=?, whatsapp_number=?, website=?, notes=?, verified_time=?, verified_by=?, statecode=?, citycode=? WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Website, u.Notes, u.VerfiedTime, u.VerifiedBy, u.Statecode, u.Citycode, u.Id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) VerifyHelpingHandLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := new(structs.ResourceLead)
	json.NewDecoder(req.Body).Decode(u)
	sqlStatement, err := uc.session.Prepare("UPDATE masterdb.helping_table SET name=?, lead_type=?, address=?, map_link=?, phone_number=?, whatsapp_number=?, website=?, notes=?, verified_time=?, verified_by=?, statecode=?, citycode=? WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(u.Name, u.LeadType, u.Address, u.URL, u.PhoneNumber, u.WhatsappNumber, u.Website, u.Notes, u.VerfiedTime, u.VerifiedBy, u.Statecode, u.Citycode, u.Id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // 201
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode(u)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) RemoveOxygenLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	sqlStatement, err := uc.session.Prepare("DELETE FROM masterdb.oxygen_table WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted) // 202
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode("hmm")
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) RemoveBedLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	sqlStatement, err := uc.session.Prepare("DELETE FROM masterdb.beds_table WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted) // 202
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode("hmm")
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) RemovePlasmaLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	sqlStatement, err := uc.session.Prepare("DELETE FROM masterdb.plasma_table WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted) // 202
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode("hmm")
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) RemoveMedicineLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	sqlStatement, err := uc.session.Prepare("DELETE FROM masterdb.medicine_table WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted) // 202
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode("hmm")
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) RemoveAmbulanceLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	sqlStatement, err := uc.session.Prepare("DELETE FROM masterdb.ambulance_table WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted) // 202
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode("hmm")
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) RemoveHelpingHandLead(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	sqlStatement, err := uc.session.Prepare("DELETE FROM masterdb.helping_table WHERE id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer sqlStatement.Close()
	tempData, err := sqlStatement.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
	rowAffected, err := tempData.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if rowAffected == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted) // 202
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError) // 500
	}
	errr := json.NewEncoder(w).Encode("hmm")
	if errr != nil {
		fmt.Println(errr)
	}
}

func (uc UserController) GetVolunteers(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	emailId := p.ByName("emailId")
	var allowed bool
	queryVolunteers, er := uc.session.Query("SELECT master FROM masterdb.loggedin WHERE email_id=?", emailId)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer queryVolunteers.Close()
	for queryVolunteers.Next() {
		err := queryVolunteers.Scan(&allowed)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if allowed != false {
		var Volunteer []*structs.Volunteer
		sqlStatement, err := uc.session.Query(" SELECT user_id, name, email_id, picture, allowed, phone_number FROM masterdb.loggedin WHERE master=? ;", false)
		if err != nil {
			panic(err)
		}
		defer sqlStatement.Close()
		for sqlStatement.Next() {
			c := new(structs.Volunteer) // initialize a new instance
			error := sqlStatement.Scan(&c.UserId, &c.Name, &c.EmailId, &c.Picture, &c.Allowed, &c.PhoneNumber)
			if error != nil {
				fmt.Println(error)
			}
			Volunteer = append(Volunteer, c) // add each instance to the slice
		}
		type temp struct {
			Volunteers []*structs.Volunteer `json:"volunteers"`
		}
		t := temp{
			Volunteers: Volunteer,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		errr := json.NewEncoder(w).Encode(t)
		if errr != nil {
			fmt.Println(errr)
		}
	} else {
		verifyResource := new(structs.VerifyResource)
		verifyResource.ResponseCode = 0
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound) // 302
		err := json.NewEncoder(w).Encode(verifyResource)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (uc UserController) GetAllResources(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	emailId := p.ByName("emailId")
	var allowed bool
	queryVolunteers, er := uc.session.Query("SELECT master FROM masterdb.loggedin WHERE email_id=?", emailId)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer queryVolunteers.Close()
	for queryVolunteers.Next() {
		err := queryVolunteers.Scan(&allowed)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if allowed != false {
		baseStruct := new(structs.ResourceMaster)
		ack := make(chan bool, 6)
		go func() {
			sqlStatement, err := uc.session.Query("SELECT * FROM masterdb.ambulance_table;")
			if err != nil {
				fmt.Println(err)
			}
			defer sqlStatement.Close()
			for sqlStatement.Next() {
				c := new(structs.ResourceLead)
				err = sqlStatement.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Ambulance = append(baseStruct.Ambulance, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.oxygen_table;")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime, &c.Id, &c.Website)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Oxygen = append(baseStruct.Oxygen, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.plasma_table;")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Plasma = append(baseStruct.Plasma, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.medicine_table;")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Medicine = append(baseStruct.Medicine, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.helping_table;")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.HelpingHand = append(baseStruct.HelpingHand, c)
			}
			ack <- true
		}()
		go func() {
			query, err := uc.session.Query("SELECT * FROM masterdb.beds_table;")
			if err != nil {
				fmt.Println(err)
			}
			defer query.Close()
			for query.Next() {
				c := new(structs.ResourceLead)
				err = query.Scan(&c.Id, &c.Website, &c.Name, &c.LeadType, &c.Statecode, &c.Citycode, &c.Address, &c.URL, &c.PhoneNumber, &c.WhatsappNumber, &c.Notes, &c.CreationTime, &c.VerifiedBy, &c.VerfiedTime)
				if err != nil {
					fmt.Println(err)
				}
				baseStruct.Bed = append(baseStruct.Bed, c)
			}
			ack <- true
		}()
		for i := 0; i < 6; i++ {
			fmt.Println(<-ack)
		}
		masterResponse := new(structs.ResourceMain)
		masterResponse.ResourceMaster = append(masterResponse.ResourceMaster, baseStruct)
		masterResponse.ResponseCode = 1
		masterResponse.CityCode = ""
		masterResponse.StateName = ""
		masterResponse.CityName = ""
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		err := json.NewEncoder(w).Encode(masterResponse)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		masterResponse := new(structs.ResourceMain)
		masterResponse.ResponseCode = 0
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound) // 302
		err := json.NewEncoder(w).Encode(masterResponse)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (uc UserController) VerifyVolunteer(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	emailId := p.ByName("emailId")
	u := structs.UpdateVolunteer{}
	json.NewDecoder(req.Body).Decode(&u)
	var allowed bool
	queryVolunteers, er := uc.session.Query("SELECT master FROM masterdb.loggedin WHERE email_id=?", emailId)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer queryVolunteers.Close()
	for queryVolunteers.Next() {
		err := queryVolunteers.Scan(&allowed)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if allowed != false {
		sqlStatement, err := uc.session.Prepare("UPDATE masterdb.loggedin SET allowed=? WHERE email_id=?;")
		if err != nil {
			fmt.Println(err)
		}
		defer sqlStatement.Close()
		tempData, err := sqlStatement.Exec(u.Allowed, u.Email)
		if err != nil {
			fmt.Println(err)
		}
		rowAffected, err := tempData.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		masterResponse := new(structs.UpdateVolunteer)
		if rowAffected == 1 {
			masterResponse.ResponseCode = 1
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // 200
		} else {
			masterResponse.ResponseCode = 0
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError) // 500

		}
		errr := json.NewEncoder(w).Encode(masterResponse)
		if errr != nil {
			fmt.Println(errr)
		}
	} else {
		masterResponse := new(structs.UpdateVolunteer)
		masterResponse.ResponseCode = 0
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		err := json.NewEncoder(w).Encode(masterResponse)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (uc UserController) RemoveVolunteer(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	emailId := p.ByName("emailId")
	u := structs.UpdateVolunteer{}
	json.NewDecoder(req.Body).Decode(&u)
	var allowed bool
	queryVolunteers, er := uc.session.Query("SELECT master FROM masterdb.loggedin WHERE email_id=?", emailId)
	if er != nil {
		fmt.Println(er)
		return
	}
	defer queryVolunteers.Close()
	for queryVolunteers.Next() {
		err := queryVolunteers.Scan(&allowed)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if allowed != false {
		sqlStatement, err := uc.session.Prepare("DELETE FROM masterdb.loggedin WHERE email_id=?;")
		if err != nil {
			fmt.Println(err)
		}
		defer sqlStatement.Close()
		tempData, err := sqlStatement.Exec(u.Email)
		if err != nil {
			fmt.Println(err)
		}
		rowAffected, err := tempData.RowsAffected()
		if err != nil {
			fmt.Println(err)
		}
		masterResponse := new(structs.UpdateVolunteer)
		if rowAffected == 1 {
			masterResponse.ResponseCode = 1
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK) // 200
		} else {
			masterResponse.ResponseCode = 0
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError) // 500

		}
		errr := json.NewEncoder(w).Encode(masterResponse)
		if errr != nil {
			fmt.Println(errr)
		}
	} else {
		masterResponse := new(structs.UpdateVolunteer)
		masterResponse.ResponseCode = 0
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		err := json.NewEncoder(w).Encode(masterResponse)
		if err != nil {
			fmt.Println(err)
		}
	}
}
