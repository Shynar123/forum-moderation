package post

import (
	"fmt"

	"forum/internal/types"
)

func (p *PostDB) CreateReport(data *types.Report) {
	_, err := p.db.Exec("INSERT INTO reports (snippet_id, snippet_title, report_type, moderator_id) VALUES ($1, $2, $3, $4)", data.PostId, data.PostTitle, data.ReportType, data.ModeratorId)
	if err != nil {
		fmt.Println("categories:", err)
		return
	}
}

func (p *PostDB) DeleteReport(postId int) {
	_, err := p.db.Exec("DELETE FROM reports WHERE snippet_id=$1", postId)
	if err != nil {
		fmt.Println("Delete post:", err)
	}
}

func (p *PostDB) ReportComplete(postId int) {
	_, err := p.db.Exec("UPDATE reports SET complete = $1 WHERE snippet_id = $2", "Complete", postId)
	if err != nil {
		fmt.Println("ReportComplete:", err)
		return
	}
}

func (p *PostDB) ReportResponse(response string, postId int) {
	_, err := p.db.Exec("UPDATE reports SET response = $1 WHERE snippet_id = $2", response, postId)
	if err != nil {
		fmt.Println("ReportResponse:", err)
		return
	}
}

func (p *PostDB) GetAllReports() []*types.Report {
	query := "SELECT * FROM reports ORDER BY id DESC"

	rows, err := p.db.Query(query)
	if err != nil {
		fmt.Println("GetAllReports1:", err)
		return nil
	}

	defer rows.Close()
	var reports []*types.Report

	for rows.Next() {
		report := types.Report{}

		err := rows.Scan(&report.Id, &report.PostId, &report.PostTitle, &report.ReportType, &report.Response)
		if err != nil {
			fmt.Println("GetAllReports2:", err)
			return nil
		}

		reports = append(reports, &report)

	}

	if err := rows.Err(); err != nil {
		fmt.Println("GetAllReports3:", err)
		return nil
	}
	return reports
}

func (p *PostDB) GetReportsByID(userId int) []*types.Report {
	query := "SELECT * FROM reports ORDER BY id DESC WHERE moderator_id=$1"

	rows, err := p.db.Query(query, userId)
	if err != nil {
		fmt.Println("GetAllReports1:", err)
		return nil
	}

	defer rows.Close()
	var reports []*types.Report

	for rows.Next() {
		report := types.Report{}

		err := rows.Scan(&report.Id, &report.PostId, &report.PostTitle, &report.ReportType, &report.Response)
		if err != nil {
			fmt.Println("GetAllReports2:", err)
			return nil
		}

		reports = append(reports, &report)

	}

	if err := rows.Err(); err != nil {
		fmt.Println("GetAllReports3:", err)
		return nil
	}
	return reports
}
