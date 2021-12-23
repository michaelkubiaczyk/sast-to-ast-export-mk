package metadata

import (
	"testing"

	"github.com/checkmarxDev/ast-sast-export/internal/app/interfaces"

	mock_app_ast_query_id "github.com/checkmarxDev/ast-sast-export/test/mocks/app/ast_query_id"
	mock_app_method_line "github.com/checkmarxDev/ast-sast-export/test/mocks/app/method_line"
	mock_app_source_file "github.com/checkmarxDev/ast-sast-export/test/mocks/app/source_file"

	mock_integration_similarity "github.com/checkmarxDev/ast-sast-export/test/mocks/integration/similarity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testResultData struct {
	MethodLines  []string
	SimilarityID string
}

func TestMetadataFactory_GetMetadataForQueryAndResult(t *testing.T) {
	astQueryID := "12532796926860742976"
	scanID := "1000001"
	metaResult1Data := testResultData{
		SimilarityID: "1234567890",
		MethodLines:  []string{"100", "2", "3", "101"},
	}
	metaResult1 := Result{
		PathID:   "2",
		ResultID: "1000002",
		FirstNode: Node{
			FileName: "Goatlin-develop/packages/clients/android/app/src/main/java/com/cx/goatlin/EditNoteActivity.kt",
			Name:     "text",
			Line:     "83",
			Column:   "78",
		},
		LastNode: Node{
			FileName: "Goatlin-develop/packages/clients/android/app/src/main/java/com/cx/goatlin/helpers/DatabaseHelper.kt",
			Name:     "note",
			Line:     "129",
			Column:   "28",
		},
	}
	metaResult2Data := testResultData{
		SimilarityID: "9492845843",
		MethodLines:  []string{"43", "21", "13"},
	}
	metaResult2 := Result{
		PathID:   "3",
		ResultID: "1000002",
		FirstNode: Node{
			FileName: "path/file1.kt",
			Name:     "text",
			Line:     "83",
			Column:   "78",
		},
		LastNode: Node{
			FileName: "path/file2.kt",
			Name:     "note",
			Line:     "129",
			Column:   "28",
		},
	}
	metaQuery := &Query{
		QueryID:  "6300",
		Language: "Kotlin",
		Name:     "SQL_Injection",
		Group:    "Kotlin_High_Risk",
		Results:  []*Result{&metaResult1, &metaResult2},
	}

	ctrl := gomock.NewController(t)
	tmpDir := t.TempDir()
	astQueryIDProviderMock := mock_app_ast_query_id.NewMockASTQueryIDRepo(ctrl)
	astQueryIDProviderMock.EXPECT().GetQueryID(metaQuery.Language, metaQuery.Name, metaQuery.Group).Return(astQueryID, nil)
	similarityIDProviderMock := mock_integration_similarity.NewMockSimilarityIDProvider(ctrl)
	similarityIDProviderMock.EXPECT().Calculate(
		gomock.Any(), metaResult1.FirstNode.Name, metaResult1.FirstNode.Line, metaResult1.FirstNode.Column, metaResult1Data.MethodLines[0],
		gomock.Any(), metaResult1.LastNode.Name, metaResult1.LastNode.Line, metaResult1.LastNode.Column, metaResult1Data.MethodLines[3],
		astQueryID,
	).Return(metaResult1Data.SimilarityID, nil)
	similarityIDProviderMock.EXPECT().Calculate(
		gomock.Any(), metaResult2.FirstNode.Name, metaResult2.FirstNode.Line, metaResult2.FirstNode.Column, metaResult2Data.MethodLines[0],
		gomock.Any(), metaResult2.LastNode.Name, metaResult2.LastNode.Line, metaResult2.LastNode.Column, metaResult2Data.MethodLines[2],
		astQueryID,
	).Return(metaResult2Data.SimilarityID, nil)
	sourceProviderMock := mock_app_source_file.NewMockSourceFileRepo(ctrl)
	sourceProviderMock.EXPECT().
		DownloadSourceFiles(scanID, gomock.Any()).
		DoAndReturn(
			func(_ string, files []interfaces.SourceFile) error {
				expectedFiles := []string{
					metaResult1.FirstNode.FileName,
					metaResult1.LastNode.FileName,
					metaResult2.FirstNode.FileName,
					metaResult2.LastNode.FileName,
				}
				var result []string
				for _, v := range files {
					result = append(result, v.RemoteName)
				}
				assert.ElementsMatch(t, expectedFiles, result)
				return nil
			},
		)
	methodLineProvider := mock_app_method_line.NewMockMethodLineRepo(ctrl)
	methodLinesResult := []*interfaces.ResultPath{
		{PathID: metaResult1.PathID, MethodLines: metaResult1Data.MethodLines},
		{PathID: metaResult2.PathID, MethodLines: metaResult2Data.MethodLines},
	}
	methodLineProvider.EXPECT().
		GetMethodLinesByPath(scanID, metaQuery.QueryID).
		Return(methodLinesResult, nil)
	metadata := NewMetadataFactory(astQueryIDProviderMock, similarityIDProviderMock, sourceProviderMock, methodLineProvider, tmpDir)

	result, err := metadata.GetMetadataRecords(scanID, metaQuery)
	assert.NoError(t, err)

	record1 := Record{
		QueryID:      metaQuery.QueryID,
		SimilarityID: metaResult1Data.SimilarityID,
		PathID:       metaResult1.PathID,
		ResultID:     metaResult1.ResultID,
	}
	record2 := Record{
		QueryID:      metaQuery.QueryID,
		SimilarityID: metaResult2Data.SimilarityID,
		PathID:       metaResult2.PathID,
		ResultID:     metaResult2.ResultID,
	}
	expectedResult := []*Record{&record1, &record2}
	assert.ElementsMatch(t, expectedResult, result)
}
